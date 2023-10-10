package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/nwoik/DRG-Mission-Monitor/api"
)

func main() {

	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/deepdives", api.DDRequestHandler)
	http.HandleFunc("/missions", api.MissionsRequestHandler)

	log.Print("Running server on port 3000")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

// Handles main page of server
func MainHandler(w http.ResponseWriter, r *http.Request) {

	// Read in the template with main webpage
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("error parsing template")

	}

	// Render the template
	t.Execute(w, nil)

	// Done.
	log.Println("Finished HTTP request at", r.URL.Path)
}
