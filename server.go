package main

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"github.com/sinfo/go-tutorial/models"
	"github.com/sinfo/go-tutorial/routes"
)

type server struct {
	mux *mux.Router
	db  *mgo.Database
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func newServer(dbURL string, dbName string) *server {
	db := models.InitDB(dbURL, dbName)
	mux := mux.NewRouter()

	server := server{mux, db}

	server.mux.HandleFunc("/hero", routes.GetHeroes).Methods("GET")
	server.mux.HandleFunc("/hero", routes.AddHero).Methods("POST")
	server.mux.HandleFunc("/hero", routes.ModifyHero).Methods("PUT")
	server.mux.HandleFunc("/hero/{id}", routes.GetHero).Methods("GET")
	server.mux.HandleFunc("/hero/{id}", routes.DeleteHero).Methods("DELETE")

	return &server
}
