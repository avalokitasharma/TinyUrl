package authservice

import (
	"authservice/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/auth/register", handler.RegisterHandler)
	http.HandleFunc("/auth/login", handler.LoginHandler)
}
