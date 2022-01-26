// Config is used for general kind of confing such as environmental variables
package config

import (
	"os"
	"strconv"
)

// GetFilePath returns the file path of where the cache should be written
// It has its default
// so even the environmental variable is not set it will work
func GetFilePath() string {
	path, ok := os.LookupEnv("DEFAULT_FILE_PATH")
	if !ok {
		path = "tmp/data.json"
	}
	return path
}

// GetFileSaveFrequency returns the frequency in seconds
// which is used as interval to save the cache to the file
// It has its default
// so even the environmental variable is not set it will work
func GetFileSaveFrequency() int {
	freq, ok := os.LookupEnv("DEFAULT_SAVING_FREQUENCY")
	if !ok {
		freq = "60"
	}

	freqINT, err := strconv.Atoi(freq)
	if err != nil {
		freqINT = 60
	}
	return freqINT
}

// GetPort returns the port number
// which is used to serve http
// It has its default
// so even the environmental variable is not set it will work
func GetPort() string {
	port, ok := os.LookupEnv("DEFAULT_PORT")
	if !ok {
		port = "3333"
	}
	return port
}
