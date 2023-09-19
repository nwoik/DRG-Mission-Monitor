package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type DeepDive struct {
	StartTime string    `json:"startTime"`
	EndTime   string    `json:"endTime"`
	Variants  []Variant `json:"variants"`
}

type Variant struct {
	Type   string  `json:"type"`
	Name   string  `json:"name"`
	Biome  string  `json:"biome"`
	Seed   int     `json:"seed"`
	Stages []Stage `json:"stages"`
}

type Stage struct {
	ID        int    `json:"id"`
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
	Anomaly   string `json:"anomaly"`
	Warning   string `json:"warming"`
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling rq")
	request, err := http.NewRequest("GET", "https://drgapi.com/v1/deepdives", nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// var body interface{}

	var properBody DeepDive

	err = json.NewDecoder(response.Body).Decode(&properBody)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(properBody)
}
