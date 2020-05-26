package app

import (
	"github.com/atomicptr/rss-merger/pkg/feed"
	"os"
	"path"
)

var feedStorage feed.Storage

func loadStorage(dir string) error {
	storageFilePath, err := getStorageFileLocation(dir)
	if err != nil {
		return err
	}

	storage, err := feed.Load(storageFilePath)
	if err != nil {
		// file isn't present, lets make a new storage...
		storage = make(feed.Storage)
	}

	feedStorage = storage
	return nil
}

func saveStorage(dir string) error {
	storageFilePath, err := getStorageFileLocation(dir)
	if err != nil {
		return err
	}

	return feed.Save(storageFilePath, feedStorage)
}

func getStorageFileLocation(dir string) (string, error) {
	if dir == "" {
		confDir, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		dir = confDir
	}

	storageDir := path.Join(dir, "rss-merger")
	err := os.MkdirAll(storageDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	return path.Join(storageDir, "storage.json"), nil
}
