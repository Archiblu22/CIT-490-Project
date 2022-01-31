package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
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
	BookId    string `json:"id"`
	BookTitle string `json:"title"`
}

func bookList(w http.ResponseWriter, r *http.Request) {
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)

	books := []Book{}

	if user.Id == 1 {
		books = append(books, Book{BookId: "1", BookTitle: "Harry Potter 1"})
		books = append(books, Book{BookId: "2", BookTitle: "Harry Potter 2"})
		books = append(books, Book{BookId: "3", BookTitle: "Harry Potter 3"})
		books = append(books, Book{BookId: "4", BookTitle: "Harry Potter 4"})
	} else {
		books = append(books, Book{BookId: "3", BookTitle: "Harry Potter 5"})
		books = append(books, Book{BookId: "4", BookTitle: "Harry Potter 6"})
	}

	json.NewEncoder(w).Encode(books)
}

func logout(w http.ResponseWriter, r *http.Request) {

}
