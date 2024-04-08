package rssreader

import (
	"github.com/mmcdole/gofeed"
	"gopkg.in/yaml.v2"
	"os"
)

type RSSFeed struct {
	URL string `yaml:"url"`
}

type Config struct {
	Feeds []RSSFeed `yaml:"feeds"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func ReadFeeds(configPath string) ([]*gofeed.Feed, error) {
	cfg, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	fp := gofeed.NewParser()
	var feeds []*gofeed.Feed
	for _, feedConfig := range cfg.Feeds {
		feed, err := fp.ParseURL(feedConfig.URL)
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, feed)
	}
	return feeds, nil
}
