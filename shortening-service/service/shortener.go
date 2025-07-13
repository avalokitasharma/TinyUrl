package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CounterResponse struct {
	ID int64 `json:"id"`
}

func ShortenURL(originalURL string, userID int) (string, error) {
	resp, err := http.Get("http://global_counter:7001/next")
	if err != nil {
		return "", fmt.Errorf("failed to get global ID: %w", err)
	}
	defer resp.Body.Close()
	var id int64
	if err := json.NewDecoder(resp.Body).Decode(&id); err != nil {
		return "", fmt.Errorf("invalid counter response: %w", err)
	}

	shortCode := base62Encode(id)
	err = repo.StoreURL(shortCode, originalURL, userID, time.Now())
	if err != nil {
		return "", fmt.Errorf("db error: %w", err)
	}
	return fmt.Sprintf("http://localhost:8080/%s", shortCode), nil

}

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func base62Encode(num uint64) string {
	if num == 0 {
		return string(charset[0])
	}
	var buf bytes.Buffer
	for num > 0 {
		rem := num % 62
		buf.WriteByte(charset[rem])
		num /= 62
	}
	// reverse
	result := buf.Bytes()
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
