package scraper

import "strings"

const (
	jsExtension   = ".js"
	cssExtension  = ".css"
	fontExtension = ".woff2"
	jpgExtension  = ".jpg"
	pngExtension  = ".png"
	jpegExtension = ".jpeg"
	gifExtension  = ".gif"
	svgExtension  = ".svg"
	icoExtension  = ".ico"
)

func isJS(url string) bool {
	return strings.Contains(url, jsExtension)
}

func isCSS(url string) bool {
	return strings.Contains(url, cssExtension)
}

func isStaticAsset(url string) bool {
	if isJS(url) || isCSS(url) || isImage(url) {
		return true
	}
	for _, ext := range []string{fontExtension} {
		if strings.Contains(url, ext) {
			return true
		}
	}
	return false
}

func isImage(link string) bool {
	for _, ext := range []string{jpegExtension, jpgExtension, pngExtension, gifExtension, svgExtension, icoExtension} {
		if strings.Contains(link, ext) {
			return true
		}
	}
	return false
}
