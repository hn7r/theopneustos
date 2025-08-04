package main

import (
	"log"
	"net/http"
)

const BaseURL = "https://api.theopneustos.bible"

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Get the passage and its decorated reference
	http.HandleFunc("/{version}/reference", handleReference)

	// Get the passage and its hypermedia
	http.HandleFunc("/{version}/text", handleText)

	log.Println("Server is running on port 7777...")
	log.Fatal(http.ListenAndServe(":7777", nil))
}
