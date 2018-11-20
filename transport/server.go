package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

/*
CreateServer
*/
func CreateServer() {

	r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
}
