package main

import (
	"github.com/drumer2142/earthquakes/handlers"
)

var feedURLs = "http://www.geophysics.geol.uoa.gr/stations/maps/seismicity.xml"

func main() {

	handlers.FetchFeed(feedURLs)
	// ticker := time.NewTicker(3 * time.Minute)
	// for range ticker.C {
	// 	handlers.FetchFeed(feedURLs)
	// }

}
