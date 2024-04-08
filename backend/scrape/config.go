package scrape

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Config represents the structure of the configuration file.
type Config struct {
	ScrapingTargets []ScrapingTarget `yaml:"scrapingTargets"`
}

// ScrapingTarget represents each scraping target in the configuration.
type ScrapingTarget struct {
	URL       string   `yaml:"url"`
	Selectors []string `yaml:"selectors"`
}

// LoadConfig loads the configuration from the given file.
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		logrus.WithError(err).Error("Failed to open config file")
		return nil, err
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		logrus.WithError(err).Error("Failed to read config file")
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse config file")
		return nil, err
	}

	return &config, nil
}
