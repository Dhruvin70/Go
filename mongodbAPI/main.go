package main

import (
	"fmt"
	"log"
	"net/http"

	routing "mongoAPI/router"
)

func main() {
	fmt.Println("MongoDB")
	r := routing.Router()
	fmt.Println("Server is getting started")
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Print("listening at port 3000....")
}
