package scrape

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/go-rod/rod"
	"github.com/sirupsen/logrus"
)

// ScrapeResult represents the result of a single scraping operation.
type ScrapeResult struct {
	URL     string
	Content map[string]string // A map of selector to scraped content
}

// ScrapeTargets performs scraping on all targets defined in the config.
// If configPath is empty, it uses a default config.yaml path.
func ScrapeTargets(configPath string) ([]ScrapeResult, error) {
	if configPath == "" {
		_, b, _, _ := runtime.Caller(0)
		basepath := filepath.Dir(b)
		configPath = filepath.Join(basepath, "config.yaml")
	}

	cfg, err := LoadConfig(configPath)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to load configuration from path: %s", configPath)
		logrus.WithError(err).Error(errMsg)
		return nil, err
	}

	var results []ScrapeResult
	for _, target := range cfg.ScrapingTargets {
		result, err := scrapeTarget(target)
		if err != nil {
			logrus.WithError(err).WithField("url", target.URL).Error("Failed to scrape target")
			continue
		}
		results = append(results, result)
	}

	return results, nil
}

// scrapeTarget handles scraping for a single target.
func scrapeTarget(target ScrapingTarget) (ScrapeResult, error) {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(target.URL)
	// Add logic to wait for necessary elements or interactions

	content := make(map[string]string)
	for _, selector := range target.Selectors {
		text, err := page.MustElement(selector).Text()
		if err != nil {
			logrus.WithError(err).WithField("selector", selector).Error("Failed to scrape content")
			continue
		}
		content[selector] = text
	}

	return ScrapeResult{URL: target.URL, Content: content}, nil
}
