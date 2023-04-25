package actualController

import (
	//coltrollers "mongoAPI/controllers"
	"encoding/json"
	controller "mongoAPI/controllers"
	modaleStruct "mongoAPI/model"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	allMovies := controller.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	var movie modaleStruct.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	controller.InsertOneMoie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "PUT")

	parmas := mux.Vars(r)
	controller.UpdateOnemovie(parmas["id"])
	json.NewEncoder(w).Encode(parmas["id"])

}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")

	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	params := mux.Vars(r)
	controller.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteallMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")

	count := controller.DeleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
