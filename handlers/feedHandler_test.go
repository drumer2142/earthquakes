package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestFetchFeed(t *testing.T) {
// 	FetchFeed(feedURLs)
// }

func TestDistance(t *testing.T) {
	fooFixedLatitude := 37.99
	fooFixedLongitude := 23.70
	fooPayload := []string{"31.8 km SSW of Levadhia", " Time: 01-Nov-2023 17:41:09 (UTC) ", " Latitude: 35.90N ", " Longitude: 28.13E ", " Depth: 2km ", " M 1.6"}
	quake := CreateQuakeData(fooPayload)
	quake.CalculateDistance(fooFixedLatitude, fooFixedLongitude, quake.Latitude, quake.Longitude, "K")
	assert.Equal(t, 457.05984795877754, quake.QuakeDistanceInKM)
}

func TestCreateQuakeData(t *testing.T) {
	fooPayload := []string{"31.8 km SSW of Levadhia", " Time: 01-Nov-2023 17:41:09 (UTC) ", " Latitude: 38.18N ", " Longitude: 22.72E ", " Depth: 2km ", " M 1.6"}
	expectedQuakeData := &QuakeData{
		Latitude:  38.18,
		Longitude: 22.72,
		Depth:     2,
		Magnitude: 1.6,
		Timestamp: "IFRpbWU6IDAxLU5vdi0yMDIzIDE3OjQxOjA5IChVVEMpIA==",
	}
	fooQuake := CreateQuakeData(fooPayload)
	assert.Equal(t, expectedQuakeData, fooQuake)
}
