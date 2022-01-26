// In memory cache system
package cache

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/canberksinangil/cache-memory/config"
)

// This is the default file permission in order to write a file.
const defaultFilePerm os.FileMode = 0666

// Cache holds the required variables to compose an in memory cache system
type Cache struct {

	// Mutex is used for handling the concurrent
	// read/write requests for cache
	mu sync.RWMutex

	// db holds the cache data
	db map[string]string
}

// NewCache creates an in memory cache system
func NewCache() *Cache {
	return &Cache{
		db: make(map[string]string),
	}
}

// GetDB returns the current cache from memory
func (c *Cache) GetDB() map[string]string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.db
}

// SyncCacheFromFile loads the key value pairs
// from given file path to the memory
// it should be called right after creating the in memory cache system
// it supports JSON format
func (c *Cache) SyncCacheFromFile() error {
	file, err := os.Open(config.GetFilePath())
	if err != nil {
		return err
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if err := json.Unmarshal(byteValue, &c.db); err != nil {
		if len(byteValue) == 0 {
			return err
		}
		return err
	}
	return nil
}

// Get reads the memory for the given key and returns its value and existence
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.db[key]
	if !ok {
		return "", false
	}
	return value, true
}

// Set will persist a value to the cache or
// override existing one with the new one
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.db[key] = value
}

// Delete deletes a given key if exists
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.db, key)

	// I could return a bool to add this info to the API response
	// _, ok := c.db[key]
	// if ok {
	// 	delete(c.db, key)
	// 	return true
	// }
	// return false
}

// Flush creates a new cache
// Right after it truncates the file
func (c *Cache) Flush() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.db = make(map[string]string)

	if err := truncateFile(); err != nil {
		return err
	}

	return nil
}

// truncateFile truncate the file to the size of 0
// which means file is emptied
func truncateFile() error {
	return os.Truncate(config.GetFilePath(), 0)
}

// StartSyncingToFile writes the cache data to the persistent file storage
// in given frequency as seconds
// it triggers the go routine
func (c *Cache) StartSyncingToFile() {
	ticker := time.NewTicker(time.Duration(config.GetFileSaveFrequency()) * time.Second)
	quit := make(chan struct{})
	go writeDataToFile(ticker, quit, c)
}

// writeDataToFile writes the data in given ticker time
func writeDataToFile(ticker *time.Ticker, quit chan struct{}, c *Cache) {
	for {
		select {
		case <-ticker.C:
			// TODO : Handle os.IsNotExist
			// TODO : Check if data has been changed

			byteValue, err := json.MarshalIndent(c.db, "", " ")
			if err != nil {
				panic(err)
			}

			if err = ioutil.WriteFile(config.GetFilePath(), byteValue, defaultFilePerm); err != nil {
				panic(err)
			} else {
				log.Println("All data added to file.")
			}

		case <-quit:
			ticker.Stop()
			return
		}
	}
}
