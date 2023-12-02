package main

import (
	"testing"

	"github.com/drumer2142/earthquakes/alerts"
)

func TestCheckLiveFeed(t *testing.T) {
	smsSender := alerts.NewSMSSender("6900000000")
	poll := alerts.NewPoller(smsSender)
	poll.Start()
}
