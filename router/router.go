package router

import (
	"github.com/gorilla/mux"
	"github.com/kushal-png/mongoapi/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteall", controller.DeleteAllMovie).Methods("DELETE")
	return r
}
