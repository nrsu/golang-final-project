package main

import (
	"log"
	"net/http"
	"bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Println("Listening 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}