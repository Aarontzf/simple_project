package main

import (
	"log"
	"net/http"

	"github.com/Aarontzf/simple_project/go-bookstore/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9091", r))
}
