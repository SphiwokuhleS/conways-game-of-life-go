package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-conways-game-of-life/domain"
	"go-conways-game-of-life/persistance"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseWorld struct {
	Name  string `json:"Name"`
	Epoch int    `json:"Epoch"`
	Grid  string `json:"Grid"`
}

type WorldHandler struct {
	Name string
}

type ResponseEpoch struct {
	Epoch int `json:"Epoch"`
}

type ResponseAllWorlds struct {
	Worlds []ResponseWorld
}

type RequestGenerationHandler struct {
	Name string      `json:"Name"`
	Grid [21][21]int `json:"Grid"`
}

type CreateWorldRequest struct {
	Name  string `json:"Name"`
	Epoch int    `json:"Epoch"`
	Grid  string `json:"Grid"`
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

func CreateWorldHandler(w http.ResponseWriter, r *http.Request) {
	var createWorldRequest CreateWorldRequest

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("Could not read create world request body", err)
	}

	reqData := json.Unmarshal(body, &CreateWorldRequest{})

	if reqData != nil {
		log.Println("Could not marshal request body", err)
	}

	world := persistance.CreateWorld(persistance.World{Name: createWorldRequest.Name, Grid: createWorldRequest.Grid, Epoch: 1})

	response := &ResponseWorld{
		Name:  world.Name,
		Grid:  world.Grid,
		Epoch: world.Epoch,
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Println("Could not marshal world created json object", err)
	}

	io.WriteString(w, string(jsonResponse))
}

func NextGeneration(w http.ResponseWriter, r *http.Request) {
	var generation RequestGenerationHandler
	var world persistance.World

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("Could not read next generation json", err)
	}
	reqData := json.Unmarshal(body, &generation)

	if reqData != nil {
		log.Println("Could not decode next generation json", reqData)
	}

	newGrid := domain.GenerateNextGeneration(generation.Grid)
	persistance.UpdateEpochWorldEpoch(generation.Name)
	persistance.UpdateGridWorldGrid(generation.Name, domain.ToString(newGrid))

	world = persistance.GetWorldByName(generation.Name)
	var response = &ResponseWorld{
		Name:  world.Name,
		Epoch: world.Epoch,
		Grid:  world.Grid,
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Println("Could not marshal world response", err)
	}

	io.WriteString(w, string(jsonResponse))
}

func GetAllWorldsHandler(w http.ResponseWriter, r *http.Request) {
	var worlds []persistance.World
	var response []ResponseWorld

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	worlds = persistance.GetAllWorlds()

	for i := 0; i < len(worlds); i++ {
		response = append(response, ResponseWorld{
			Name:  worlds[i].Name,
			Epoch: worlds[i].Epoch,
			Grid:  worlds[i].Grid,
		})
	}

	jsonWorlds, _ := json.Marshal(response)
	io.WriteString(w, string(jsonWorlds))
}

func GetWorldEpoch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	name, err := vars["name"]

	if !err {
		log.Println(err)
	}

	worldEpoch := persistance.GetWorldEpoch(name)

	var response = ResponseEpoch{
		Epoch: worldEpoch,
	}

	jsonEpoch, _ := json.Marshal(response)

	io.WriteString(w, string(jsonEpoch))
}

func GetWorldHandler(w http.ResponseWriter, r *http.Request) {
	var world persistance.World
	var worldRequest WorldHandler

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	} else {
		req := json.Unmarshal(body, &worldRequest)
		if req != nil {
			log.Println(req)
			return
		}
	}

	world = persistance.GetWorldByName(worldRequest.Name)
	var response = &ResponseWorld{
		Name:  world.Name,
		Epoch: world.Epoch,
		Grid:  world.Grid,
	}

	jsonWorld, _ := json.Marshal(response)
	io.WriteString(w, string(jsonWorld))
}

func DeleteWorldHandler() {
	//
}
