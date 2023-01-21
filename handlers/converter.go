package handlers

import (
	"encoding/json"
	"github.com/drumer2142/earthquakes/types"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	ActivityCounter = 3
)

type Filter struct {
	filter *types.Filters
}

type QuakeData struct {
	Latitude  string
	Longitude string
	Depth     string
	Magnitude string
}

func feedsConverter(feeds []*types.GeophysicsRss) {
	for _, geo := range feeds {
		for i := 0; i < len(geo.Channel.Items); i++ {
			if i > ActivityCounter {
				break
			}
			descriptionItems := strings.Split(geo.Channel.Items[i].Description, "<br>")
			log.Println(descriptionItems)

			quake := cleanData(descriptionItems)
			quake.filterActivity()
		}

	}
}

func cleanData(item []string) *QuakeData {
	floatReg, _ := regexp.Compile("[+-]?([0-9]*[.])?[0-9]+")

	return &QuakeData{
		Latitude:  floatReg.FindString(item[2]),
		Longitude: floatReg.FindString(item[3]),
		Depth:     floatReg.FindString(item[4]),
		Magnitude: floatReg.FindString(item[5]),
	}
}

func (quake *QuakeData) filterActivity() {
	filters := loadFilters()
	log.Println(filters)
	//for _, filter := range filters.Parameters {
	//	if filter.Magnitude ==  {
	//
	//	}
	//}
}

func loadFilters() types.Filters {
	jsonFile, err := os.Open("filters.json")
	defer jsonFile.Close()

	if err != nil {
		log.Println(err)
	}
	var filters types.Filters
	jsonParser := json.NewDecoder(jsonFile)
	err = jsonParser.Decode(&filters)

	if err != nil {
		log.Fatalln(err)
	}

	return filters
}
