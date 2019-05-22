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

	staticFilesHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))

	mux.PathPrefix("/static/").Handler(staticFilesHandler)

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	// swagger:route GET /hero heroes GetHeroes
	//
	// Returns all heroes
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	mux.HandleFunc("/hero", routes.GetHeroes).Methods("GET")

	// swagger:route POST /hero heroes AddHero
	//
	// Returns new hero
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	mux.HandleFunc("/hero", routes.AddHero).Methods("POST")

	// swagger:route PUT /hero heroes ModifyHero
	//
	// Returns modified hero
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	mux.HandleFunc("/hero", routes.ModifyHero).Methods("PUT")

	// swagger:route GET /hero/{id} heroes GetHero
	//
	// Returns a specific hero
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	mux.HandleFunc("/hero/{id}", routes.GetHero).Methods("GET")

	// swagger:route DELETE /hero/{id} heroes DeleteHero
	//
	// Returns the deleted hero
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	mux.HandleFunc("/hero/{id}", routes.DeleteHero).Methods("DELETE")

	ServerInstance = &Server{handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(mux), db}
}
