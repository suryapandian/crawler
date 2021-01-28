package scraper

import (
	"fmt"
	"net/url"
	"strings"
)

func isRelativeLink(link string) bool {
	linkURI, err := url.Parse(link)
	if err != nil {
		return false
	}
	return !linkURI.IsAbs()
}

func isInternalLink(link, current string) bool {
	currentURI, _ := url.Parse(current)
	return strings.Contains(link, currentURI.Host)
}

func expandLink(link, current string) string {
	linkURI, _ := url.Parse(link)
	base, _ := url.Parse(current)
	return fmt.Sprintf("%v", base.ResolveReference(linkURI))
}
