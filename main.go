package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

	http.HandleFunc("/login", login)
	http.HandleFunc("/logged-in", loggedIn)
	http.HandleFunc("/logout", logout)

}

func login(w http.ResponseWriter, r *http.Request) {

}

func loggedIn(w http.ResponseWriter, r *http.Request) {

}

func logout(w http.ResponseWriter, r *http.Request) {

}
