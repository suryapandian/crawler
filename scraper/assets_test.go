package scraper

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type AssetTestSuite struct {
	suite.Suite
}

func TestAssetTestSuite(t *testing.T) {
	suite.Run(t, new(AssetTestSuite))
}

func (t *AssetTestSuite) TestIsJS() {
	t.True(isJS("https://widget.trustpilot.com/bootstrap/v5/tp.widget.bootstrap.min.js"))

	t.False(isJS("https://use.typekit.net/xcc2pjd.css"))
}

func (t *AssetTestSuite) TestIsCSS() {
	t.False(isCSS("https://widget.trustpilot.com/bootstrap/v5/tp.widget.bootstrap.min.js"))

	t.True(isCSS("https://use.typekit.net/xcc2pjd.css"))
}

func (t *AssetTestSuite) TestIsImage() {
	var testCases = []struct {
		imageURL string
		isValid  bool
	}{
		{
			"https://content.cuvva.com/segments/home_header.png",
			true,
		},
		{
			"https://i.redd.it/519u0np73xb61.jpg",
			true,
		},
		{
			"https://use.typekit.net/xcc2pjd.css",
			false,
		},
		{
			"https://widget.trustpilot.com/bootstrap/v5/tp.widget.bootstrap.min.js",
			false,
		},
	}
	for _, tc := range testCases {
		isValidImageURL := isImage(tc.imageURL)
		t.Equal(isValidImageURL, tc.isValid)
	}

}
