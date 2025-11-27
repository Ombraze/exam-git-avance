package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Serveur Go en démarrage...")


	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/book", BookHandler)
	http.HandleFunc("/contact", ContactHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bonjour depuis HomeHandler")
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bonjour depuis BookHandler")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bonjour depuis ContactHandler")
}
