package worker

import (
	"context"
	"sync"

	"github.com/suryapandian/crawler/logger"
	"github.com/suryapandian/crawler/scraper"
)

/*To Do:
make a generic worker with input and output channel with input and output function
let the worker accept the max number of workers that can be run, the size of the input and output channels
*/

type Crawler struct {
	urlsToCrawl  chan string
	siteMaps     chan scraper.SiteMap
	crawlerCache *CrawlerCache
	maxScraper   int
	url          string
	ctx          context.Context
	cancel       context.CancelFunc
}

func NewCrawler(url string) *Crawler {
	var crawledSites sync.Map
	return &Crawler{
		urlsToCrawl:  make(chan string, 10),
		siteMaps:     make(chan scraper.SiteMap, 1000),
		crawlerCache: &CrawlerCache{sites: &crawledSites},
		maxScraper:   5,
		url:          url,
	}
}

func (c *Crawler) Run() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.urlsToCrawl <- c.url
	for i := 1; i <= c.maxScraper; i++ {
		go scrape(c.urlsToCrawl, c.siteMaps, c.crawlerCache, c.ctx)
	}
	c.getSiteMap()
}

func scrape(urls chan string, sitemaps chan scraper.SiteMap, cache *CrawlerCache, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case url := <-urls:
			if cache.hasCrawled(url) { //Skip scraping again, the channel could have duplicates
				continue
			}

			siteMap, err := scraper.NewScraper(url, logger.LogEntryWithRef()).Scrape()
			if err != nil {
				return
			}

			if hasCrawled := cache.save(url, siteMap); !hasCrawled {
				sitemaps <- siteMap
			}
		}
	}

}

func (c *Crawler) getSiteMap() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case siteMap := <-c.siteMaps:
			siteMap.PrettyPrint()
			go c.getSiteMapOfInternalLinks(siteMap.InternalLinks)
		}
	}
}

func (c *Crawler) getSiteMapOfInternalLinks(links []string) {
	select {
	case <-c.ctx.Done():
		close(c.urlsToCrawl)
	default:
		for _, link := range links {
			if c.crawlerCache.hasCrawled(link) {
				continue
			}
			c.urlsToCrawl <- link
		}
	}
}

func (c *Crawler) Stop() {
	logger.LogEntryWithRef().Infoln("Gracelly stopping the crawler")
	c.cancel()
}
