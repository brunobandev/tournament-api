package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Tournament :)
type Tournament struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string        `json:"name,omitempty"`
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/tournaments", TournamentList).Methods("GET")
	router.HandleFunc("/tournament/create", TournamentCreate).Methods("POST")

	log.Fatal(http.ListenAndServe(":4001", router))
}

func TournamentList(w http.ResponseWriter, r *http.Request) {

	var tournaments []Tournament

	session, _ := mgo.Dial("172.17.0.2")
	defer session.Close()

	err := session.DB("local").C("tournaments").Find(nil).All(&tournaments)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(tournaments)
}

func TournamentCreate(w http.ResponseWriter, r *http.Request) {

	session, _ := mgo.Dial("172.17.0.2")
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("local").C("tournaments")

	decoder := json.NewDecoder(r.Body)
	var t Tournament

	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	_ = c.Insert(t)

	// tournaments = append(tournaments, t)

	json.NewEncoder(w).Encode(t)
}
