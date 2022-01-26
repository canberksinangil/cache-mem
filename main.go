package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/canberksinangil/cache-memory/api"
	"github.com/canberksinangil/cache-memory/cache"
	"github.com/canberksinangil/cache-memory/config"
)

func main() {
	cache := cache.NewCache()

	if err := cache.SyncCacheFromFile(); err != nil {
		log.Printf("There is been error while syncing the file. Error is: %s", err.Error())
	}

	cache.StartSyncingToFile()
	log.Printf("Initial cache is %v", cache.GetDB())

	ch := api.NewCacheHandler(cache)

	http.HandleFunc("/healthz", ch.HealthCheck)
	http.HandleFunc("/cache", ch.Cache)
	http.HandleFunc("/flush", ch.Flush)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.GetPort()), api.ServerLogger(http.DefaultServeMux)); err != nil {
		log.Fatal(err)
	}
}
