package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-conways-game-of-life/api"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", api.PingHandler).Methods("GET", "POST", "PUT", "DELETE")
	r.HandleFunc("/world", api.GetWorldHandler).Methods("POST")
	r.HandleFunc("/worlds", api.GetAllWorldsHandler).Methods("GET")
	r.HandleFunc("/world/{name}", api.GetWorldEpoch).Methods("GET")
	r.HandleFunc("/create", api.CreateWorldHandler).Methods("POST")
	r.HandleFunc("/generation", api.NextGeneration).Methods("POST")
	r.HandleFunc("/generations", api.MultipleGenerations).Methods("POST")
	r.HandleFunc("/world/{name}", api.DeleteWorldHandler).Methods("DELETE")

	fmt.Println("Running game of life server..")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Println("Couldn't start server", err)
		return
	}
}
