package main

import (
	"TasksAPI/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port string = "8080"

func main() {
	log.Println("Start server")
	router := mux.NewRouter()

	utils.BuildTaskRoutes(router)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
