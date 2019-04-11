package server

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"github.com/sinfo/go-tutorial/models"
	"github.com/sinfo/go-tutorial/routes"
)

type Server struct {
	Mux *mux.Router
	DB  *mgo.Database
}

var ServerInstance *Server

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}

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
