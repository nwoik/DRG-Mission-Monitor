package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type MissonsResponse struct {
	Biomes Biomes `json:"Biomes"`
}

type Biomes struct {
	CrystallineCaverns       []Stage `json:"Crystalline Caverns"`
	DenseBiozone             []Stage `json:"Dense Biozone"`
	FungosBogs               []Stage `json:"Fungos Bogs"`
	MagmaCore                []Stage `json:"Magma Core"`
	RadioactiveExclusionZone []Stage `json:"Radioactive Exclusion Zone"`
	SaltPits                 []Stage `json:"Salt Pits"`
}

type DeepDiveResponse struct {
	DeepDive DeepDiveData `json:"Deep Dives"`
}

type DeepDiveData struct {
	DeepDiveElite  DeepDive `json:"Deep Dive Elite"`
	DeepDiveNormal DeepDive `json:"Deep Dive Normal"`
}

type DeepDive struct {
	Biome    string  `json:"Biome"`
	CodeName string  `json:"CodeName"`
	Stages   []Stage `json:"Stages"`
}

type Stage struct {
	CodeName           string   `json:"CodeName"`
	Complexity         string   `json:"Complexity"`
	Length             string   `json:"Length"`
	ID                 int      `json:"id"`
	PrimaryObjective   string   `json:"PrimaryObjective"`
	SecondaryObjective string   `json:"SecondaryObjective"`
	MissionMutator     string   `json:"MissionMutator"`
	MissionWarnings    []string `json:"MissionWarnings"`
}

func DDRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling rq")
	request, err := http.NewRequest("GET", "https://doublexp.net/json?data=DD", nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// var body interface{}

	// err = json.NewDecoder(response.Body).Decode(&body)
	// if err != nil {
	// 	log.Print(err)
	// }
	// log.Println(body)

	var properBody DeepDiveResponse

	err = json.NewDecoder(response.Body).Decode(&properBody)
	if err != nil {
		log.Print(err)
	}
	log.Println(properBody)
}

func MissionsRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling rq")
	request, err := http.NewRequest("GET", "https://doublexp.net/json?data=current", nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// var body interface{}

	var properBody MissonsResponse

	err = json.NewDecoder(response.Body).Decode(&properBody)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(properBody)
}
