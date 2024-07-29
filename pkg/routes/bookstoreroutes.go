package routes

import (
	"github.com/gorilla/mux"
	"github.com/moaz/go-BookStoreProject/pkg/controllers"
)

var RegisterBookroutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
