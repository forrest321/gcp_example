package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/forrest321/gcp_example/rssreader" // Adjust the import path as necessary
)

func main() {
	configPath := flag.String("config", filepath.Join("backend", "rssreader", "config.yaml"), "Path to the configuration file")
	flag.Parse()

	feeds, err := rssreader.ReadFeeds(*configPath)
	if err != nil {
		log.Fatalf("Error reading RSS feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed: %s\n", feed.Title)
		for _, item := range feed.Items {
			fmt.Printf("Title: %s\nLink: %s\n\n", item.Title, item.Link)
		}
	}
}
