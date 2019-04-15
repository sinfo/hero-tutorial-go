package server

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"github.com/sinfo/go-tutorial/models"
	"github.com/sinfo/go-tutorial/routes"
)

// Server representes the server, and stores its router and database instance
type Server struct {
	Mux *mux.Router
	DB  *mgo.Database
}

// ServerInstance represents the instance of the running server
var ServerInstance *Server

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}

// InitServer initializes the server and its database

func InitServer(dbURL string, dbName string) {
	db := models.InitDB(dbURL, dbName)
	mux := mux.NewRouter()

	ServerInstance = &Server{mux, db}

	ServerInstance.Mux.HandleFunc("/hero", routes.GetHeroes).Methods("GET")
	ServerInstance.Mux.HandleFunc("/hero", routes.AddHero).Methods("POST")
	ServerInstance.Mux.HandleFunc("/hero", routes.ModifyHero).Methods("PUT")
	ServerInstance.Mux.HandleFunc("/hero/{id}", routes.GetHero).Methods("GET")
	ServerInstance.Mux.HandleFunc("/hero/{id}", routes.DeleteHero).Methods("DELETE")

}
