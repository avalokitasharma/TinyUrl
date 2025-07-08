package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/shorten", proxyToShorteningService)
	http.HandleFunc("/", proxyToRedirectService)

	log.Println("API gateway listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func proxyToShorteningService(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://shortening_service:7002/shorten", http.StatusTemporaryRedirect)
}

func proxyToRedirectService(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://redirect_service:7003"+r.URL.Path, http.StatusTemporaryRedirect)
}
