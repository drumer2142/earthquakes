package handlers

import (
	"encoding/json"
	"github.com/drumer2142/earthquakes/types"
	"log"
	"os"
	"strings"
)

var (
	ActivityCounter = 3
)

type Filter struct {
	filter *types.Filters
}

func feedsConverter(feeds []*types.GeophysicsRss) {
	for _, geo := range feeds {
		for i := 0; i < len(geo.Channel.Items); i++ {
			if i > ActivityCounter {
				break
			}
			descriptionItems := strings.Split(geo.Channel.Items[i].Description, "<br>")
			log.Println(descriptionItems)

			filterActivity(descriptionItems)
			//distanceFromAthens := descriptionItem[0]
			//timeOfEvent := descriptionItem[1]
			//lat := descriptionItem[2]
			//long := descriptionItem[3]
			//depth := descriptionItem[4]
			//magnitude := descriptionItem[5]
		}

	}
}

func filterActivity(items []string) {
	loadFilters()
}

func loadFilters() {
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
}
