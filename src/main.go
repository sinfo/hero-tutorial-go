package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
)

const (
	URL        = "localhost:27017"
	Database   = "gotutorial"
	Collection = "heroes"
)

type Hero struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

// mock
var heroes []Hero

var session *mgo.Session

func GetHeroes(w http.ResponseWriter, r *http.Request) {
	var heroes []Hero
	c := session.DB(Database).C(Collection)
	err := c.Find(nil).All(&heroes)

	if err == nil {
		http.Error(w, "Unable to make query do database", http.StatusExpectationFailed)
	}

	json.NewEncoder(w).Encode(heroes)
}

func AddHero(w http.ResponseWriter, r *http.Request) {

}

func GetHero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err == nil {
		fmt.Println(id)
	}

	for _, item := range heroes {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

func ModifyHero(w http.ResponseWriter, r *http.Request) {}
func DeleteHero(w http.ResponseWriter, r *http.Request) {}

func main() {
	// hero mock setup
	var err error

	session, err = mgo.Dial(URL)

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/hero", GetHeroes).Methods("GET")
	router.HandleFunc("/hero", AddHero).Methods("POST")
	router.HandleFunc("/hero/{id}", GetHero).Methods("GET")
	router.HandleFunc("/hero/{id}", ModifyHero).Methods("PUT")
	router.HandleFunc("/hero/{id}", DeleteHero).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
