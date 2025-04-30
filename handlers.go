package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	QueryTypeRange     = "RANGE"
	QueryTypeVerseID   = "ID"
	QueryTypeReference = "REF"
)

func handleReference(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetReference(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleText(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetText(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetReference(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var query = r.URL.Query().Get("q")
	version := strings.ToUpper(r.PathValue("version"))

	if !isValidVersion(version) {
		http.Error(w, "Invalid version", http.StatusConflict)
		return
	}

	enc := json.NewEncoder(w)
	reference := parseQuery(query)
	if !isValidReference(reference) {
		http.Error(w, "Invalid reference", http.StatusConflict)
		return
	}

	passage, err := getPassage(reference.book, reference.chapter, reference.verse, version)
	if err != nil {
		http.Error(w, "Failed locating resource", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"reference": decorateReference(reference),
		"text":      passage,
	}

	err = enc.Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	return
}

func GetText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var query = r.URL.Query().Get("q")
	version := strings.ToUpper(r.PathValue("version"))

	if !isValidVersion(version) {
		http.Error(w, "Invalid version", http.StatusConflict)
		return
	}

	enc := json.NewEncoder(w)
	reference := parseQuery(query)
	if !isValidReference(reference) {
		http.Error(w, "Invalid reference", http.StatusConflict)
		return
	}

	passage, err := getPassage(reference.book, reference.chapter, reference.verse, version)
	if err != nil {
		http.Error(w, "Failed locating resource", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"reference": decorateReference(reference),
		"text":      passage,
		//"prevVerse": getPreviousVerse(reference),
		//"nextVerse": getNextVerse(reference),
	}

	err = enc.Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	return
}
