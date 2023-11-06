package main

import (
	"log"
	"time"

	"github.com/drumer2142/earthquakes/handlers"
)

var (
	feedURLs     = "http://www.geophysics.geol.uoa.gr/stations/maps/seismicity.xml"
	pollInterval = 3
)

func main() {

	handlers.FetchFeed(feedURLs)
	ticker := time.NewTicker(time.Duration(pollInterval) * time.Minute)

	for {
		err := handlers.FetchFeed(feedURLs)
		if err != nil {
			log.Fatal(err)
		}
		<-ticker.C
	}

}
