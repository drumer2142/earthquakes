package handlers

import (
	"encoding/xml"
	"github.com/drumer2142/earthquakes/types"
	"io"
	"log"
	"net/http"
	"time"
)

type feedResult struct {
	feed *types.GeophysicsRss
	err  error
}

func FetchFeeds(urls []string) {
	feedChan := make(chan feedResult, len(urls))

	for _, url := range urls {
		go fetchSingleFeed(url, feedChan)
	}

	var feeds []*types.GeophysicsRss

	for i := 0; i < len(urls); i++ {
		res := <-feedChan
		// If the goroutine errors out, we'll just wait for other feeds
		if res.err != nil {
			continue
		}
		feeds = append(feeds, res.feed)

		feedsConverter(feeds)
	}
}

func fetchSingleFeed(url string, feedChan chan feedResult) {
	net := &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := net.Get(url)

	log.Println(res.Body)
	// If there was an error write that to the channel return
	if err != nil {
		feedChan <- feedResult{nil, err}
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			feedChan <- feedResult{nil, err}
			return
		}
	}(res.Body)

	feed, err := parseFeed(res.Body)

	if err != nil {
		feedChan <- feedResult{nil, err}
		return
	}

	feedChan <- feedResult{feed, nil}

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
