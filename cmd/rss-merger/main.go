package main

import (
	"github.com/atomicptr/rss-merger/pkg/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
