package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiteshshimpi-55/go-mongo-api/controller"
)

func Router() *mux.Router {

	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./static")))
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")          // Get all Movies
	r.HandleFunc("/movie/{id}", controller.GetOneMovies).Methods("GET")      // Get one Movie
	r.HandleFunc("/movie", controller.CreateMovie).Methods("POST")           // Create and Post new Movie
	r.HandleFunc("/movie/{id}", controller.Update).Methods("PUT")            // Update one Movie
	r.HandleFunc("/movie/{id}", controller.DeleteOneMovie).Methods("DELETE") // Delete one Movie
	r.HandleFunc("/movies", controller.DeleteAllMovies).Methods("DELETE")    // Delete all Movies

	return r
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Everyone</h1>"))
}
