package scraper

import (
	"golang.org/x/net/html"
)

const (
	anchorTag   = "a"
	scriptTag   = "script"
	linkTag     = "link"
	imageTag    = "img"
	hrefTagKey  = "href"
	assetTagKey = "src"
)

func getTagValue(t html.Token, attributeKey string) (ok bool, val string) {
	for _, a := range t.Attr {
		if a.Key == attributeKey {
			val = a.Val
			ok = true
		}
	}
	return
}
