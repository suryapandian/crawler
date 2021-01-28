package scraper

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suryapandian/crawler/logger"
)

func TestScraper(t *testing.T) {
	a := assert.New(t)

	file, err := os.Open("../test/test1.html")
	if err != nil {
		t.Fatalf("loading file failed: %v", err)
	}
	defer file.Close()
	body, err := ioutil.ReadAll(file)
	a.NoError(err)

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(body)
	}))
	defer mockServer.Close()
	serverURL := mockServer.URL

	expectedStaticAssets := []string{
		serverURL + "/lib/l119.js",
		serverURL + "/locale/en_US/duckduckgo14.js",
		serverURL + "/util/u523.js",
		serverURL + "/d2902.js",
		serverURL + "/s1949.css",
		serverURL + "/o1949.css",
		serverURL + "/font/ProximaNova-Reg-webfont.woff2",
		serverURL + "/font/ProximaNova-Sbold-webfont.woff2",
		serverURL + "/font/ProximaNova-ExtraBold-webfont.woff2",
		serverURL + "/assets/icons/meta/DDG-iOS-icon_152x152.png",
		serverURL + "/assets/icons/meta/DDG-iOS-icon_120x120.png",
		serverURL + "/assets/icons/meta/DDG-icon_256x256.png",
		serverURL + "/assets/icons/meta/DDG-iOS-icon_76x76.png",
		serverURL + "/assets/icons/meta/DDG-iOS-icon_60x60.png",
		serverURL + "/manifest.json",
		serverURL + "/favicon.ico",
	}

	expectedInternalLinks := []string{
		serverURL + "/about",
	}

	expectedExternalLinks := []string{
		"https://www.freeformatter.com",
		"https://duckduckgo.com/",
	}

	siteMap, err := NewScraper(serverURL, logger.LogEntryWithRef()).Scrape()
	a.NoError(err)
	a.NotEmpty(siteMap)

	//Check siteMap
	a.ElementsMatch(expectedStaticAssets, siteMap.StaticAssets)
	a.ElementsMatch(expectedInternalLinks, siteMap.InternalLinks)
	a.ElementsMatch(expectedExternalLinks, siteMap.ExternalLinks)
}

func TestScraperError(t *testing.T) {
	a := assert.New(t)

	siteMap, err := NewScraper("invalidURL", logger.LogEntryWithRef()).Scrape()
	a.Error(err)
	a.Empty(siteMap)
}
