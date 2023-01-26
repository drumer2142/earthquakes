package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestFetchFeedWithTicker(t *testing.T) {
	for i := 0; i < 4; i++ {
		FetchFeeds(testURLs)
	}
}

func TestDistance(t *testing.T) {
	lat := 35.90
	long := 28.13
	quakeDistanceInKM := distance(FixedLatitude, FixedLongitude, lat, long, "K")
	assert.Equal(t, 457.05984795877754, quakeDistanceInKM)
}
