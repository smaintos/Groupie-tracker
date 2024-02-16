package main

import (
	back "Track/back"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", back.Artist)
	http.HandleFunc("/artiste", back.Detail)
	http.HandleFunc("/search", back.Search)
	http.HandleFunc("/erreur", back.Error)
	http.ListenAndServe(":8080", nil)

}
