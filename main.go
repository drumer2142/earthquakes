package main

import (
	"github.com/drumer2142/earthquakes/handlers"
	"time"
)

var feedURLs = []string{
	"http://www.geophysics.geol.uoa.gr/stations/maps/seismicity.xml",
}

func main() {
	ticker := time.NewTicker(3 * time.Minute)
	for _ = range ticker.C {
		handlers.FetchFeeds(feedURLs)
	}

}
