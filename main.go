package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Tournament :)
type Tournament struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

var tournaments []Tournament

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tournament/create", TournamentCreate).Methods("POST")

	log.Fatal(http.ListenAndServe(":4001", router))
}

func TournamentCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t Tournament

	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	tournaments = append(tournaments, t)

	json.NewEncoder(w).Encode(tournaments)
}
