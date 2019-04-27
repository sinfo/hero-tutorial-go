package server

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/sinfo/hero-tutorial-go/src/models"
	"github.com/sinfo/hero-tutorial-go/src/routes"
)

// Server representes the server, and stores its router and database instance
type Server struct {
	Mux http.Handler
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

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	mux.HandleFunc("/hero", routes.GetHeroes).Methods("GET")
	mux.HandleFunc("/hero", routes.AddHero).Methods("POST")
	mux.HandleFunc("/hero", routes.ModifyHero).Methods("PUT")
	mux.HandleFunc("/hero/{id}", routes.GetHero).Methods("GET")
	mux.HandleFunc("/hero/{id}", routes.DeleteHero).Methods("DELETE")

	mux.HandleFunc("/swagger", routes.GetSwagger).Methods("GET")

	ServerInstance = &Server{handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(mux), db}
}
