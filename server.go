package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Pages/index.html")
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Pages/project.html")
}

func teamHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Pages/teams.html")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/project", projectHandler)
	http.HandleFunc("/team", teamHandler)
	http.Handle("/assetts/", http.StripPrefix("/assetts/", http.FileServer(http.Dir("assetts"))))
	fmt.Println("Listening on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
