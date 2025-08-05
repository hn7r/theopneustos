package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
)

var ApiKeys []string

func main() {
	startup()
	registerApiKeys()

	fs := http.FileServer(http.Dir("public"))
	mux := http.NewServeMux()

	mux.Handle("/", fs)
	mux.Handle("/{version}/reference", apiKeyMiddleware(http.HandlerFunc(handleReference)))
	mux.Handle("/{version}/text", apiKeyMiddleware(http.HandlerFunc(handleText)))

	log.Fatal(http.ListenAndServe(":7777", mux))
}

func registerApiKeys() {
	ApiKeys = append(ApiKeys, os.Getenv("API_KEY_HOMELAB_GO"))
	ApiKeys = append(ApiKeys, os.Getenv("API_KEY_HOMELAB_GO_2"))
}

func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "no auth", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(token, " ")
		if len(parts) != 2 {
			http.Error(w, "invalid auth token", http.StatusUnauthorized)
			return
		}

		splitKey := parts[1]
		if !slices.Contains(ApiKeys, splitKey) {
			http.Error(w, "invalid auth token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func startup() {
	banner := `
░██████████░██     ░██ ░██████████   ░██████   ░█████████  ░███    ░██ ░██████████ ░██     ░██   ░██████   ░██████████  ░██████     ░██████   
    ░██    ░██     ░██ ░██          ░██   ░██  ░██     ░██ ░████   ░██ ░██         ░██     ░██  ░██   ░██      ░██     ░██   ░██   ░██   ░██  
    ░██    ░██     ░██ ░██         ░██     ░██ ░██     ░██ ░██░██  ░██ ░██         ░██     ░██ ░██             ░██    ░██     ░██ ░██         
    ░██    ░██████████ ░█████████  ░██     ░██ ░█████████  ░██ ░██ ░██ ░█████████  ░██     ░██  ░████████      ░██    ░██     ░██  ░████████  
    ░██    ░██     ░██ ░██         ░██     ░██ ░██         ░██  ░██░██ ░██         ░██     ░██         ░██     ░██    ░██     ░██         ░██ 
    ░██    ░██     ░██ ░██          ░██   ░██  ░██         ░██   ░████ ░██          ░██   ░██   ░██   ░██      ░██     ░██   ░██   ░██   ░██  
    ░██    ░██     ░██ ░██████████   ░██████   ░██         ░██    ░███ ░██████████   ░██████     ░██████       ░██      ░██████     ░██████
`
	fmt.Print(banner)
	fmt.Println("Listening on port 7777...")
}
