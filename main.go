package main

import (
	"log"
	"main/server"
	"net/http"
)

func main() {
	log.Println("Serveur Go en démarrage...")

	http.HandleFunc("/", server.HomeHandler)
	http.HandleFunc("/book", server.BookHandler)
	http.HandleFunc("/contact", server.ContactHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
