package main

import (
	"fmt"
	"log"
	"net/http"

	groupie "groupie/Handlers"
)

func main() {
	http.HandleFunc("/", groupie.HomeHandler)
	http.HandleFunc("/artist/", groupie.ArtistHandler)
	// now handle
	http.HandleFunc("/style/", groupie.StyleHandler)

	fmt.Println("Server starting on http://localhost:8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
