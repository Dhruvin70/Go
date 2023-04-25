package routing

import (
	actualController "mongoAPI/Actualcountollers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", actualController.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/createmovie", actualController.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", actualController.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/deletemovie/{id}", actualController.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", actualController.DeleteallMovie).Methods("DELETE")

	return router
}
