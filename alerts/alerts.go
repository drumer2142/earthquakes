package alerts

import (
	"log"
	"time"

	"github.com/drumer2142/earthquakes/handlers"
	"github.com/drumer2142/earthquakes/types"
)

var (
	feedURLs        = "http://www.geophysics.geol.uoa.gr/stations/maps/seismicity.xml"
	pollInterval    = 1 * time.Minute
	FixedLatitude   = 37.99
	FixedLongitude  = 23.70
	NotifTimestamps = []string{"test"}
)

type Sender interface {
	Send(*handlers.QuakeData) error
}

type Poller struct {
	Sender
}

type SMSSender struct {
	number string
}

type PushNotification struct {
	channel string
}

func NewPoller(sender Sender) *Poller {
	return &Poller{
		Sender: sender,
	}
}

func NewSMSSender(number string) *SMSSender {
	return &SMSSender{
		number: number,
	}
}

func NewPushNotification(channel string) *PushNotification {
	return &PushNotification{
		channel: channel,
	}
}

func (poller *Poller) Start() {
	ticker := time.NewTicker(time.Duration(pollInterval))

	for {
		quakesMap, err := handlers.FetchFeed(feedURLs)
		if err != nil {
			log.Fatal(err)
		}
		<-ticker.C

		for _, quake := range quakesMap {
			poller.SendAlert(quake)
		}
	}
}

func (poller *Poller) SendAlert(quake *handlers.QuakeData) {
	if filterActivity(quake) {
		poller.Sender.Send(quake)
	}
}

func (s *SMSSender) Send(quake *handlers.QuakeData) error {
	log.Printf("SEND QUAKE ALERT MG=%f DST=%f DEPTH=%f \n", quake.Magnitude, quake.QuakeDistanceInKM, quake.Depth)
	log.Println("Sending Alert to number: ", s.number)
	return nil
}

func filterActivity(quake *handlers.QuakeData) bool {
	filters := handlers.LoadFilters()
	quake.CalculateDistance(FixedLatitude, FixedLongitude, quake.Latitude, quake.Longitude, "K")
	log.Println("Distance Of Quake In KM:", quake.QuakeDistanceInKM)
	for _, filter := range filters.Parameters {

		if filterCriteriaAreMet(filter, quake) {
			if !checkDuplicatesExist(quake) {
				return true
			}
		}
	}
	return false
}

func filterCriteriaAreMet(filter types.Parameters, quake *handlers.QuakeData) bool {
	if quake.Magnitude >= filter.MinMagnitude && quake.QuakeDistanceInKM <= filter.MaxDistanseInKM && quake.Depth >= filter.MinDepth {
		return true
	}
	return false
}

func checkDuplicatesExist(quake *handlers.QuakeData) bool {
	for _, stamp := range NotifTimestamps {
		if quake.Timestamp != stamp {
			NotifTimestamps = append(NotifTimestamps, quake.Timestamp)
			return false
		}
	}
	return true
}
