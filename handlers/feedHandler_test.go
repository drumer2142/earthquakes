package handlers

import (
	"testing"

	"github.com/drumer2142/earthquakes/types"
	"github.com/stretchr/testify/assert"
)

var feedURLs = "http://www.geophysics.geol.uoa.gr/stations/maps/seismicity.xml"

func TestFetchFeed(t *testing.T) {
	FetchFeed(feedURLs)
}

func TestDistance(t *testing.T) {
	lat := 35.90
	long := 28.13
	quakeDistanceInKM := calculateDistance(FixedLatitude, FixedLongitude, lat, long, "K")
	assert.Equal(t, 457.05984795877754, quakeDistanceInKM)
}

func TestCreateQuakeData(t *testing.T) {
	descriptionItems := []string{"31.8 km SSW of Levadhia", " Time: 01-Nov-2023 17:41:09 (UTC) ", " Latitude: 38.18N ", " Longitude: 22.72E ", " Depth: 2km ", " M 1.6"}

	expectedQuakeData := &QuakeData{
		Latitude:  38.18,
		Longitude: 22.72,
		Depth:     2,
		Magnitude: 1.6,
		Timestamp: "IFRpbWU6IDAxLU5vdi0yMDIzIDE3OjQxOjA5IChVVEMpIA==",
	}

	quake := createQuakeData(descriptionItems)

	assert.Equal(t, expectedQuakeData, quake)
}

func TestFilterCriteriaAreMet(t *testing.T) {
	quakeDistanceInKM := 80.0

	filter1 := types.Parameters{
		MaxDistanseInKM: 64,
		Timestamp:       "",
		MinDepth:        3,
		MinMagnitude:    2.1,
	}

	descriptionItems := []string{"29.3 km NNW of Chalkida", "Time: 03-Nov-2023 06:26:38 (UTC)", "Latitude: 38.72N", "Longitude: 23.53E", "Depth: 6km", "M 5.2"}
	quake := createQuakeData(descriptionItems)
	status1 := quake.filterCriteriaAreMet(filter1, quakeDistanceInKM)

	filter2 := types.Parameters{
		MaxDistanseInKM: 81,
		Timestamp:       "",
		MinDepth:        3,
		MinMagnitude:    1.8,
	}

	descriptionItems = []string{"29.3 km NNW of Chalkida", "Time: 03-Nov-2023 06:26:38 (UTC)", "Latitude: 38.72N", "Longitude: 23.53E", "Depth: 10km", "M 2.0"}
	quake = createQuakeData(descriptionItems)
	status2 := quake.filterCriteriaAreMet(filter2, quakeDistanceInKM)

	assert.Equal(t, false, status1)
	assert.Equal(t, true, status2)
}
