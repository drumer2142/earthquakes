package alerts

import (
	"testing"

	"github.com/drumer2142/earthquakes/handlers"
	"github.com/drumer2142/earthquakes/types"
	"github.com/stretchr/testify/assert"
)

func TestFilterCriteriaAreMet(t *testing.T) {
	fooFixedLatitude := 37.99
	fooFixedLongitude := 23.70

	filter1 := types.Parameters{
		MaxDistanseInKM: 60,
		Timestamp:       "",
		MinDepth:        3,
		MinMagnitude:    2.1,
	}

	descriptionItems := []string{"29.3 km NNW of Chalkida", "Time: 03-Nov-2023 06:26:38 (UTC)", "Latitude: 38.72N", "Longitude: 23.53E", "Depth: 6km", "M 5.2"}
	quake := handlers.CreateQuakeData(descriptionItems)
	quake.CalculateDistance(fooFixedLatitude, fooFixedLongitude, quake.Latitude, quake.Longitude, "K")
	status1 := filterCriteriaAreMet(filter1, quake)

	filter2 := types.Parameters{
		MaxDistanseInKM: 83,
		Timestamp:       "",
		MinDepth:        3,
		MinMagnitude:    1.8,
	}

	descriptionItems = []string{"29.3 km NNW of Chalkida", "Time: 03-Nov-2023 06:26:38 (UTC)", "Latitude: 38.72N", "Longitude: 23.53E", "Depth: 10km", "M 2.0"}
	quake = handlers.CreateQuakeData(descriptionItems)
	quake.CalculateDistance(fooFixedLatitude, fooFixedLongitude, quake.Latitude, quake.Longitude, "K")
	status2 := filterCriteriaAreMet(filter2, quake)

	assert.Equal(t, false, status1)
	assert.Equal(t, true, status2)
}
