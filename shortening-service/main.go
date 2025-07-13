package shorteningservice

import (
	"log"
	"net/http"
	"shorteningservice/middleware"
)

func main() {
	http.Handle("/shorten", middleware.AuthMiddleware(http.HandleFunc(handler.ShortenHandler)))
	log.Println("Shortening service running on: 7002")
	http.ListenAndServe(":7002", nil)
}
