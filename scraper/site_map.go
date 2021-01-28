package scraper

import (
	"fmt"
	"github.com/suryapandian/crawler/array"
)

type SiteMap struct {
	URL           string
	StaticAssets  []string
	InternalLinks []string
	ExternalLinks []string
}

func (w *SiteMap) PrettyPrint() {
	fmt.Printf("\n \n \n \t List of static assets that %s depends on: \n \n \n ", w.URL)
	prettyPrintStringArray(w.StaticAssets)
}

func (s *SiteMap) clean() {
	s.StaticAssets = array.RemoveDuplicates(s.StaticAssets)
	s.InternalLinks = array.RemoveDuplicates(s.InternalLinks)
	s.ExternalLinks = array.RemoveDuplicates(s.ExternalLinks)
}

func prettyPrintStringArray(arr []string) {
	for _, v := range arr {
		fmt.Printf("\t \t %s \n", v)
	}
}
