package handlers

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	"github.com/drumer2142/earthquakes/types"
)

func FetchFeed(url string) error {

	net := &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := net.Get(url)

	if err != nil {
		return err
	}

	feed, err := parseFeed(res.Body)

	if err != nil {
		return err
	}

	feedsConverter(feed)

	return err
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
