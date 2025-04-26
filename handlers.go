package main

import (
	"encoding/json"
	"net/http"
)

func handleReference(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetText(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetText(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query().Get("q")
	var version = r.URL.Query().Get("v")

	w.Header().Set("Content-Type", "application/json")

	if !validateVersion(version) {
		http.Error(w, "Invalid version", http.StatusConflict)
		return
	}

	if !validateQuery(query) {
		http.Error(w, "Invalid reference ID", http.StatusConflict)
		return
	}

	enc := json.NewEncoder(w)
	reference := parseQuery(query)
	passage, err := getPassage(reference.book, reference.chapter, reference.verse, version)
	if err != nil {
		http.Error(w, "Failed locating resource", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"text":      passage,
		"reference": decorateReference(reference),
	}

	err = enc.Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	
	return
}
