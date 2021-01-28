package worker

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suryapandian/crawler/scraper"
)

func TestCrawler(t *testing.T) {
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

	urlsToCrawl := make(chan string, 10)
	siteMaps := make(chan scraper.SiteMap, 100)
	var crawledSites sync.Map
	crawlerCache := &CrawlerCache{sites: &crawledSites}
	ctx, cancel := context.WithCancel(context.Background())

	a.Empty(siteMaps)
	urlsToCrawl <- serverURL
	go scrape(urlsToCrawl, siteMaps, crawlerCache, ctx)
	time.Sleep(2 * time.Second)
	cancel()

	a.NotEmpty(siteMaps)
	a.Len(siteMaps, 1)
	siteMap := <-siteMaps
	a.NotEmpty(siteMap.InternalLinks)
	a.NotEmpty(siteMap.StaticAssets)
	a.NotEmpty(siteMap.ExternalLinks)

	//Trying to push the same URL again and asserting that siteMaps channel does not have duplicated sitemap
	urlsToCrawl <- serverURL
	go scrape(urlsToCrawl, siteMaps, crawlerCache, ctx)
	time.Sleep(2 * time.Second)
	cancel()

	a.Empty(siteMaps)
}
