package repo

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://user:pass@db:5432/tinyurl?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
func CreateUser(username, password string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, hash)
	return err
}

func ValidateUser(username, password string) (struct{ ID int }, error) {
	var id int
	var hash string
	err := db.QueryRow("SELECT id, password_hash FROM users WHERE username=$1", username).Scan(&id, &hash)
	if err != nil {
		return struct{ ID int }{}, err
	}
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return struct{ ID int }{}, sql.ErrNoRows
	}
	return struct{ ID int }{ID: id}, nil
}
