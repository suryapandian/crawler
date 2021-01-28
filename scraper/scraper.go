package scraper

import (
	"golang.org/x/net/html"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Scraper struct {
	url     string
	logger  *logrus.Entry
	SiteMap SiteMap
}

func NewScraper(url string, logger *logrus.Entry) *Scraper {
	return &Scraper{
		url:    url,
		logger: logger,
	}
}

func (s *Scraper) Scrape() (SiteMap, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		s.logger.Errorf("error while scraping url: %v", err)
		return s.SiteMap, err
	}
	s.SiteMap.URL = s.url

	s.parse(resp)
	s.SiteMap.clean()
	return s.SiteMap, nil
}

func (s *Scraper) parse(resp *http.Response) {

	body := resp.Body
	defer body.Close()

	tokenizer := html.NewTokenizer(body)

	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case (tt == html.StartTagToken) || (tt == html.SelfClosingTagToken):
			t := tokenizer.Token()
			switch t.Data {
			case anchorTag:
				s.parseHrefTag(t)
			case scriptTag:
				s.parseScriptTag(t)
			case linkTag:
				s.parseLinkTag(t)
			case imageTag:
				s.parseImageTag(t)
			default:
				//s.logger.Infoln("Default", t.Data)
			}

		}
	}

}

func (s *Scraper) parseHrefTag(t html.Token) {
	ok, url := getTagValue(t, hrefTagKey)
	if !ok {
		return
	}
	s.saveLink(url)
}

func (s *Scraper) saveLink(url string) {
	switch {
	case isRelativeLink(url):
		s.SiteMap.InternalLinks = append(s.SiteMap.InternalLinks, expandLink(url, s.url))
	case isInternalLink(url, s.url):
		s.SiteMap.InternalLinks = append(s.SiteMap.InternalLinks, url)
	default:
		s.SiteMap.ExternalLinks = append(s.SiteMap.ExternalLinks, url)
	}
}

func (s *Scraper) parseScriptTag(t html.Token) {
	ok, src := getTagValue(t, assetTagKey)
	if !ok {
		return
	}
	if isRelativeLink(src) {
		src = expandLink(src, s.url)
	}

	if isJS(src) || isCSS(src) {
		s.SiteMap.StaticAssets = append(s.SiteMap.StaticAssets, src)
		return
	}
	s.logger.Infoln("script", src)

}

func (s *Scraper) parseLinkTag(t html.Token) {
	ok, src := getTagValue(t, hrefTagKey)
	if !ok {
		return
	}

	if isRelativeLink(src) {
		src = expandLink(src, s.url)
	}

	if isStaticAsset(src) {
		s.SiteMap.StaticAssets = append(s.SiteMap.StaticAssets, src)
		return
	}

	s.saveLink(src)
}

func (s *Scraper) parseImageTag(t html.Token) {
	ok, src := getTagValue(t, assetTagKey)
	if !ok {
		return
	}
	s.SiteMap.StaticAssets = append(s.SiteMap.StaticAssets, src)
}
