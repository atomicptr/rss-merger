package app

import (
	"crypto/md5"
	"fmt"
	"github.com/atomicptr/rss-merger/pkg/rss"
	"github.com/mmcdole/gofeed"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"io"
	"log"
	"time"
)

func createMergedFeed(identifier string) (interface{}, error) {
	cacheIdentifier := fmt.Sprintf("rss:%s", identifier)
	content, ok := appCache.Get(cacheIdentifier)
	if ok {
		return content, nil
	}

	log.Printf("No cache found for identifier \"%s\"...", identifier)

	f, ok := feedStorage[identifier]
	if !ok {
		return nil, fmt.Errorf("unknown identifier: %s", identifier)
	}

	feedParser := gofeed.NewParser()

	newFeed := rss.New(f.Title, "Created by github.com/atomicptr/rss-merger", siteLink)
	newFeed.Channel.LastBuildDate = time.Now().Format(time.RFC822)

	for _, link := range f.Links {
		parsedFeed, err := feedParser.ParseURL(link)
		if err != nil {
			log.Println(errors.Wrapf(err, "error with link '%s'", link))
			continue
		}

		for _, parsedFeedItem := range parsedFeed.Items {
			if parsedFeedItem.Link == "" {
				log.Println("error: feed has invalid link")
				continue
			}

			guid := parsedFeedItem.GUID

			if guid == "" {
				guid = makeGuid(parsedFeedItem.Link)
			}

			item := rss.Item{
				Title:   parsedFeedItem.Title,
				Link:    parsedFeedItem.Link,
				GUID:    guid,
				PubDate: parsedFeedItem.Published,
			}

			newFeed.Item = append(newFeed.Item, item)
		}
	}

	appCache.Set(cacheIdentifier, newFeed, cache.DefaultExpiration)
	return newFeed, nil
}

func makeGuid(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
