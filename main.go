package main

import (
	"github.com/drumer2142/earthquakes/handlers"
	"log"
)

var feedURLs = []string{
	"http://www.geophysics.geol.uoa.gr/stations/maps/seismicity.xml",
}

func main() {
	feeds := handlers.FetchFeeds(feedURLs)
	log.Println(feeds)
}
