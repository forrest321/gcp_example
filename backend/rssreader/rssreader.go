// File: backend/rssreader/rssreader.go

package rssreader

import (
	"github.com/mmcdole/gofeed"
	// Include other necessary imports
)

// RSSFeed represents a single RSS feed with its URL and optional selectors.
type RSSFeed struct {
	URL string `yaml:"url"`
}

// Config represents the structure of the configuration file for RSS feeds.
type Config struct {
	Feeds []RSSFeed `yaml:"feeds"`
}

// LoadConfig loads the RSS feed configuration from the given file.
// Similar to the previous implementation in scraper.go
func LoadConfig(filename string) (*Config, error) {
	// Implementation...
}

// ReadFeeds reads RSS feeds from the provided list of feed URLs or from the config file.
func ReadFeeds(feedURLs []string, configPath string) ([]*gofeed.Feed, error) {
	var feeds []*gofeed.Feed
	var urls []string

	if len(feedURLs) > 0 {
		urls = feedURLs
	} else {
		cfg, err := LoadConfig(configPath)
		if err != nil {
			return nil, err
		}
		for _, feed := range cfg.Feeds {
			urls = append(urls, feed.URL)
		}
	}

	fp := gofeed.NewParser()
	for _, url := range urls {
		feed, err := fp.ParseURL(url)
		if err != nil {
			// Handle or log the error
			continue
		}
		feeds = append(feeds, feed)
	}

	return feeds, nil
}
