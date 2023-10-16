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
	AzureWorld               *[]Stage `json:"Azure World,omitempty"`
	CrystallineCaverns       *[]Stage `json:"Crystalline Caverns,omitempty"`
	DenseBiozone             *[]Stage `json:"Dense Biozone,omitempty"`
	FungusBogs               *[]Stage `json:"Fungus Bogs,omitempty"`
	GlacialStrata            *[]Stage `json:"Glacial Strata,omitempty"`
	HollowBough              *[]Stage `json:"Hollow Bough,omitempty"`
	MagmaCore                *[]Stage `json:"Magma Core,omitempty"`
	RadioactiveExclusionZone *[]Stage `json:"Radioactive Exclusion Zone,omitempty"`
	SaltPits                 *[]Stage `json:"Salt Pits,omitempty"`
	SandblastedCorridors     *[]Stage `json:"Sandblasted Corridors,omitempty"`
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
	PrimaryObjective   string    `json:"PrimaryObjective"`
	SecondaryObjective string    `json:"SecondaryObjective"`
	CodeName           string    `json:"CodeName"`
	Complexity         string    `json:"Complexity"`
	Length             string    `json:"Length"`
	ID                 int       `json:"id"`
	MissionMutator     *string   `json:"MissionMutator,omitempty"`
	MissionWarnings    *[]string `json:"MissionWarnings,omitempty"`
}

func DDRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling rq")
	request, err := http.NewRequest("GET", "https://doublexp.net/json?data=DD", nil)
	if err != nil {
		log.Print(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Print(err)
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
	EnableCors(&w)

	log.Println("Handling rq")
	request, err := http.NewRequest("GET", "https://doublexp.net/json?data=current", nil)
	if err != nil {
		log.Print(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// var body interface{}

	var properBody MissonsResponse

	err = json.NewDecoder(response.Body).Decode(&properBody)
	if err != nil {
		log.Print(err)
	}
	log.Println(properBody)

	resp, err := json.Marshal(&properBody)
	if err != nil {
		log.Print(err)
	}

	w.Write(resp)
	log.Print("response written")
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Push-Key, Push-Token")
	(*w).Header().Set("Content-Type", "Application/json")
}
