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
	Name  string `json:"name"`
	Epoch int    `json:"epoch"`
	Grid  string `json:"grid"`
}

type WorldHandler struct {
	Name string
}

type ResponseEpoch struct {
	Epoch int `json:"epoch"`
}

type ResponseAllWorlds struct {
	Worlds []ResponseWorld
}

type RequestGenerationHandler struct {
	Name string `json:"name"`
	Grid string `json:"grid"`
}

type CreateWorldRequest struct {
	Name  string `json:"name"`
	Epoch int    `json:"epoch"`
	Grid  string `json:"grid"`
}

type MultiGenerationRequest struct {
	Name        string `json:"name"`
	Grid        string `json:"grid"`
	Generations int    `json:"generations"`
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

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("Could not read create world request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqData := json.Unmarshal(body, &createWorldRequest)

	if reqData != nil {
		log.Println("Could not marshal request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
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
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, string(jsonResponse))
	return
}

func NextGeneration(w http.ResponseWriter, r *http.Request) {
	var generation RequestGenerationHandler
	var world persistance.World
	var grid [21][21]int

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

	err = json.Unmarshal([]byte(generation.Grid), &grid)
	if err != nil {
		log.Println("Could not unmarshal array", err)
	}

	newGrid := domain.GenerateNextGeneration(grid)
	newGridJson := domain.ToJsonArray(domain.ToString(newGrid))
	persistance.UpdateEpochWorldEpoch(generation.Name)
	persistance.UpdateGridWorldGrid(generation.Name, newGridJson)

	world = persistance.GetWorldByName(generation.Name)
	gridResponse := domain.ToJsonArray(world.Grid)
	var response = &ResponseWorld{
		Name:  world.Name,
		Epoch: world.Epoch,
		Grid:  gridResponse,
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Println("Could not marshal world response", err)
	}

	io.WriteString(w, string(jsonResponse))
}

func MultipleGenerations(w http.ResponseWriter, r *http.Request) {
	var generation MultiGenerationRequest
	var world persistance.World
	var grid [21][21]int

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

	err = json.Unmarshal([]byte(generation.Grid), &grid)
	if err != nil {
		log.Println("Could not unmarshal array", err)
	}

	newGrid := domain.MultipleGenerations(generation.Generations, grid)
	newGridJson := domain.ToJsonArray(domain.ToString(newGrid))
	persistance.UpdateEpochWorldEpoch(generation.Name)
	persistance.UpdateGridWorldGrid(generation.Name, newGridJson)

	world = persistance.GetWorldByName(generation.Name)
	gridResponse := domain.ToJsonArray(world.Grid)
	var response = &ResponseWorld{
		Name:  world.Name,
		Epoch: world.Epoch,
		Grid:  gridResponse,
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

func DeleteWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	name, err := vars["name"]

	if !err {
		log.Println(err)
	}

	world := persistance.GetWorldByName(name)
	deleted := persistance.DeleteWorld(world)

	if deleted {
		io.WriteString(w, `{"deleted": true, "message": "World deleted"}`)
	} else {
		io.WriteString(w, `{"deleted": false, "message":"World does not exist"}`)
	}
}
