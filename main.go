package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/sub2nwoik", MainHandler)

	// Assuming there is a server.crt and server.key file existing in the local directory, run TLS server
	log.Print("Running server on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
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
