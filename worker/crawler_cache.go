package worker

import (
	"sync"

	"github.com/suryapandian/crawler/scraper"
)

type CrawlerCache struct {
	sites *sync.Map
}

//ToDo - based on last updatedAt tag

func (c *CrawlerCache) save(site string, sitemap scraper.SiteMap) bool {
	_, crawled := c.sites.LoadOrStore(site, sitemap)
	return crawled
}

func (c *CrawlerCache) hasCrawled(site string) bool {
	_, ok := c.sites.Load(site)
	return ok
}
