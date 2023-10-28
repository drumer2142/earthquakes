package handlers

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/drumer2142/earthquakes/types"
)

type feedResult struct {
	feed *types.GeophysicsRss
	err  error
}

func FetchFeed(url string) {

	net := &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := net.Get(url)

	if err != nil {
		log.Println(err)
	}

	feed, err := parseFeed(res.Body)

	if err != nil {
		log.Println(err)
	}

	feedsConverter(feed)

	log.Println(res.Body)

}

func parseFeed(body io.ReadCloser) (*types.GeophysicsRss, error) {
	var geophysicsRss types.GeophysicsRss

	decoder := xml.NewDecoder(body)
	err := decoder.Decode(&geophysicsRss)
	if err != nil {
		return nil, err
	}

	return &geophysicsRss, nil
}
