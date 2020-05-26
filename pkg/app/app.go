package app

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var appCache *cache.Cache

var storageDir = ""
var siteLink = ""

func Run() error {
	storageDir = getEnv("RSS_MERGER_STORAGEDIR", "")
	siteLink = getEnv("RSS_MERGER_SITELINK", "")
	cacheDuration := time.Duration(getEnvAsInt("RSS_MERGER_CACHE_DURATION_IN_MINUTES", 5))

	username := getEnv("RSS_MERGER_USERNAME", "")
	password := getEnv("RSS_MERGER_PASSWORD", "")
	port := getEnvAsInt("RSS_MERGER_PORT", 8081)

	err := loadStorage(storageDir)
	if err != nil {
		return err
	}

	appCache = cache.New(cacheDuration*time.Minute, 15*time.Minute)
	return runApi(username, password, port)
}
