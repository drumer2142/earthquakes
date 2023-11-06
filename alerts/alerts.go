package alerts

import (
	"log"
	"time"

	"github.com/drumer2142/earthquakes/handlers"
)

var (
	feedURLs     = "http://www.geophysics.geol.uoa.gr/stations/maps/seismicity.xml"
	pollInterval = 1 * time.Minute
)

type SMSSender struct {
	number string
}

type Sender interface {
	Send(*handlers.QuakeData) error
}

type Poller struct {
	Sender
}

func NewPoller(sender Sender) *Poller {
	return &Poller{
		Sender: sender,
	}
}

func (poller *Poller) Start() {
	ticker := time.NewTicker(time.Duration(pollInterval))

	for {
		err := handlers.FetchFeed(feedURLs)
		if err != nil {
			log.Fatal(err)
		}
		<-ticker.C
	}
}

func (poller *Poller) SendAlert(alert *handlers.QuakeData) error {
	return poller.Sender.Send(alert)
}

func NewSMSSender(number string) *SMSSender {
	return &SMSSender{
		number: number,
	}
}

func (s *SMSSender) Send(alert *handlers.QuakeData) error {
	log.Println("Sending Alert to number: ", s.number)
	return nil
}
