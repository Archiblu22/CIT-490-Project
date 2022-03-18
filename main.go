package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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
	port := os.Getenv("PORT")

	dbx, err = sqlx.Open("sqlite3", "database/books.db")
	handleError(err)
	_, err = dbx.Exec("SELECT 1")
	handleError(err)

	http.HandleFunc("/", serveTemplate)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/login", login)
	http.HandleFunc("/book-list", bookList)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/add-to-library", addToLibrary)
	http.HandleFunc("/remove-from-library", removeFromLibrary)
	http.HandleFunc("/signup", addNewUser)
	http.HandleFunc("/logged-in", loggedIn)

	log.Println("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

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

	users := []User{}
	err := dbx.Select(&users, "SELECT id FROM user WHERE email = ? AND password = ?", login.Email, login.Password)
	handleError(err)
	if len(users) > 0 {
		json.NewEncoder(w).Encode(LoginResult{
			Id:       users[0].Id,
			LoggedIn: true,
		})
	} else {
		json.NewEncoder(w).Encode(LoginResult{
			Id:       0,
			LoggedIn: false,
		})
	}
}

func loggedIn(w http.ResponseWriter, r *http.Request) {
}

type User struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type Book struct {
	UserId       int    `json:"userid" db:"userid"`
	GoogleBookId string `json:"googlebookid" db:"googlebookid"`
}

func bookList(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"][0]
	userid, err := strconv.Atoi(id)
	handleError(err)
	books := []Book{}
	err = dbx.Select(&books, "SELECT userid, googlebookid FROM library WHERE userid = ?", userid)
	handleError(err)
	json.NewEncoder(w).Encode(books)
}

func logout(w http.ResponseWriter, r *http.Request) {

}

func addToLibrary(w http.ResponseWriter, r *http.Request) {
	book := Book{}
	json.NewDecoder(r.Body).Decode(&book)
	handleError(err)
	books := []Book{}
	err = dbx.Select(&books, "SELECT userid, googlebookid FROM library WHERE userid = ? AND googlebookid = ?", book.UserId, book.GoogleBookId)

	if len(books) > 0 {
		return
	}

	_, err := dbx.Exec("INSERT INTO library (userid, googlebookid) VALUES (?, ?)", book.UserId, book.GoogleBookId)
	handleError(err)
}

func removeFromLibrary(w http.ResponseWriter, r *http.Request) {
	book := Book{}
	json.NewDecoder(r.Body).Decode(&book)
	_, err := dbx.Exec("DELETE FROM library WHERE userid = ? AND googlebookid = ?", book.UserId, book.GoogleBookId)
	handleError(err)
}

func addNewUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	_, err := dbx.Exec("INSERT INTO user (email, password) VALUES (?, ?)", user.Email, user.Password)
	handleError(err)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("www", "layout.html")
	fp := filepath.Join("www", filepath.Clean(r.URL.Path))
	if r.URL.Path == "/" {
		fp = filepath.Join("www", filepath.Clean("/index.html"))
	}

	// Return a 404 if the template doesn't exist or it's a directory
	info, err := os.Stat(fp)
	if err != nil || info.IsDir() {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	tmpl, err := template.New(r.URL.Path).Delims("((", "))").ParseFiles(lp, fp)
	handleError(err)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	handleError(err)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}
