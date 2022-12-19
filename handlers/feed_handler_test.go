package handlers

import "testing"

var testURLs = []string{
	"http://www.geophysics.geol.uoa.gr/stations/maps/seismicity.xml",
}

func TestFetchFeed(t *testing.T) {
	FetchFeeds(testURLs)
}

func TestFetchSingleFeed(t *testing.T) {
	for _, url := range testURLs {
		feedChan := make(chan feedResult, len(testURLs))
		fetchSingleFeed(url, feedChan)
	}

}
