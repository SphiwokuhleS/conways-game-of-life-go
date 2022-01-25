package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-conways-game-of-life/api"
	"net/http"
	"strings"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", api.PingHandler).Methods("GET", "POST", "PUT", "DELETE")
	r.HandleFunc("/world", api.GetWorldHandler).Methods("POST")
	r.HandleFunc("/worlds", api.GetAllWorldsHandler).Methods("GET")
	r.HandleFunc("/world/{name}", api.GetWorldEpoch).Methods("GET")

	fmt.Println("Running game of life server")
	http.ListenAndServe(":8080", r)

}

/**
        Convert 2D array to string so it can be stored in the database as json
**/
func toString(grid [10][10]int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(grid)), " "), "")
}
