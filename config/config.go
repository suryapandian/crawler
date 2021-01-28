package config

import (
	"os"
)

var (
	URL_TO_CRAWL string
	LOG_LEVEL    string
)

func init() {
	URL_TO_CRAWL = os.Getenv("URL_TO_CRAWL")
	if URL_TO_CRAWL == "" {
		URL_TO_CRAWL = "https://www.cuvva.com/"
	}
	LOG_LEVEL = os.Getenv("LOG_LEVEL")
	if LOG_LEVEL == "" {
		LOG_LEVEL = "INFO"
	}
}
