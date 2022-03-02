package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var dbx *sqlx.DB
var err error

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	dbx, err = sqlx.Open("sqlite3", "database/books.db")
	handleError(err)
	_, err = dbx.Exec("SELECT 1")
	handleError(err)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/login", login)
	http.HandleFunc("/book-list", bookList)
	http.HandleFunc("/logout", logout)

	log.Println("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	Id       int  `json:"id"`
	LoggedIn bool `json:"loggedIn"`
}

func login(w http.ResponseWriter, r *http.Request) {
	login := Login{}
	json.NewDecoder(r.Body).Decode(&login)
	fmt.Printf("Email: %v, Password: %v\n", login.Email, login.Password)

	id := 0

	if login.Email == "pierre@pierre.com" {
		id = 1
	} else {
		id = 2
	}

	loggedIn := login.Password == "12345"

	json.NewEncoder(w).Encode(LoginResult{
		Id:       id,
		LoggedIn: loggedIn,
	})
}

func loggedIn(w http.ResponseWriter, r *http.Request) {
}

type User struct {
	Id int `json:"id"`
}

type Book struct {
	UserId       string `json:"id" db:"user_id"`
	GoogleBookId string `json:"google_book_id" db:"google_book_id"`
	Date         string `json:"date" db:"date"`
}

func bookList(w http.ResponseWriter, r *http.Request) {
	books := []Book{}
	err := dbx.Select(&books, "SELECT user_id, google_book_id, date FROM library")
	handleError(err)
	json.NewEncoder(w).Encode(books)
}

func logout(w http.ResponseWriter, r *http.Request) {

}
