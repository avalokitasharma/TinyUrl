package handler

import (
	"encoding/json"
	"net/http"
	"shorteningservice/middleware"
)

type Request struct {
	OriginalURL string `json:"original_url`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	shortURL, err := service.ShortenURL(req.OriginalURL, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
}
