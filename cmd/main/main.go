package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moaz/go-BookStoreProject/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookroutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}
