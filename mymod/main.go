package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hellow mod in golang")
	greeter()
}

func greeter() {
	fmt.Println("hi how r u")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}

// w is for write responce back
// r is for somebody sends responce

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcone to go lang</h1>"))

}
