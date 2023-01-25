package handlers

import (
	"encoding/json"
	"github.com/drumer2142/earthquakes/types"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	ActivityCounter = 20
	FixedLatitude   = 37.99
	FixedLongitude  = 23.70
)

type Filter struct {
	filter *types.Filters
}

type QuakeData struct {
	Latitude  float64
	Longitude float64
	Depth     float64
	Magnitude float64
}

func feedsConverter(feeds []*types.GeophysicsRss) {
	for _, geo := range feeds {
		for i := 0; i < len(geo.Channel.Items); i++ {
			if i > ActivityCounter {
				break
			}
			descriptionItems := strings.Split(geo.Channel.Items[i].Description, "<br>")
			//log.Println("Description Items: ", descriptionItems)

			quake := createQuakeData(descriptionItems)
			quake.filterActivity()
		}

	}
}

func createQuakeData(item []string) *QuakeData {
	floatReg, _ := regexp.Compile("[+-]?([0-9]*[.])?[0-9]+")

	lat, _ := strconv.ParseFloat(floatReg.FindString(item[2]), 64)
	long, _ := strconv.ParseFloat(floatReg.FindString(item[3]), 64)
	depth, _ := strconv.ParseFloat(floatReg.FindString(item[4]), 64)
	magnitude, _ := strconv.ParseFloat(floatReg.FindString(item[5]), 64)

	return &QuakeData{
		Latitude:  lat,
		Longitude: long,
		Depth:     depth,
		Magnitude: magnitude,
	}
}

func (quake *QuakeData) filterActivity() {
	filters := loadFilters()
	quakeDistanceInKM := distance(FixedLatitude, FixedLongitude, quake.Latitude, quake.Longitude, "K")
	for _, filter := range filters.Parameters {

		if quake.Magnitude >= filter.Magnitude && quakeDistanceInKM <= filter.DistanceInKm && quake.Depth >= filter.Depth {
			log.Printf("SEND QUAKE ALERT MG=%f DST=%f DEPTH=%f", quake.Magnitude, quakeDistanceInKM, quake.Depth)
		}
	}

}

func loadFilters() types.Filters {
	jsonFile, err := os.Open("handlers/filters.json")
	defer jsonFile.Close()

	if err != nil {
		log.Println("err: ", err)
	}
	var filters types.Filters
	jsonParser := json.NewDecoder(jsonFile)
	err = jsonParser.Decode(&filters)

	if err != nil {
		log.Println("err: ", err)
	}

	return filters
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//:::                                                                         :::
//:::  This routine calculates the distance between two points (given the     :::
//:::  latitude/longitude of those points). It is being used to calculate     :::
//:::  the distance between two locations using GeoDataSource (TM) products  :::
//:::                                                                         :::
//:::  Definitions:                                                           :::
//:::    South latitudes are negative, east longitudes are positive           :::
//:::                                                                         :::
//:::  Passed to function:                                                    :::
//:::    lat1, lon1 = Latitude and Longitude of point 1 (in decimal degrees)  :::
//:::    lat2, lon2 = Latitude and Longitude of point 2 (in decimal degrees)  :::
//:::    unit = the unit you desire for results                               :::
//:::           where: 'M' is statute miles (default)                         :::
//:::                  'K' is kilometers                                      :::
//:::                  'N' is nautical miles                                  :::
//:::                                                                         :::
//:::  Worldwide cities and other features databases with latitude longitude  :::
//:::  are available at https://www.geodatasource.com                         :::
//:::                                                                         :::
//:::  For enquiries, please contact sales@geodatasource.com                  :::
//:::                                                                         :::
//:::  Official Web site: https://www.geodatasource.com                       :::
//:::                                                                         :::
//:::               GeoDataSource.com (C) All Rights Reserved 2022            :::
//:::                                                                         :::
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}
