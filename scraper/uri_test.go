package scraper

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type URITestSuite struct {
	suite.Suite
}

func TestURITestSuite(t *testing.T) {
	suite.Run(t, new(URITestSuite))
}

func (t *URITestSuite) TestIsRelativeLink() {
	t.True(isRelativeLink("/home"))

	t.False(isRelativeLink("https://cuvva.com/careers"))
}

func (t *URITestSuite) TestIsInternalLink() {
	t.True(isInternalLink("https://cuvva.com/careers", "https://cuvva.com/home"))

	t.False(isInternalLink("https://use.joy.net/sessions", "https://cuvva.com/home"))
}

func (t *URITestSuite) TestExpandLink() {
	var testCases = []struct {
		relativeLink string
		currentLink  string
		expectedLink string
	}{
		{
			"/home",
			"https://cuvva.com/careers",
			"https://cuvva.com/home",
		},
		{
			"/products",
			"https://cuvva.com/careers",
			"https://cuvva.com/products",
		},
	}
	for _, tc := range testCases {
		expandedLink := expandLink(tc.relativeLink, tc.currentLink)
		t.Equal(expandedLink, tc.expectedLink)
	}

}
