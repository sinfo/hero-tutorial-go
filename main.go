package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sinfo/go-tutorial/models"
	"github.com/sinfo/go-tutorial/routes"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Configuration struct {
	Host string
	Port int
	DB   *mgo.Database
}

var Config *Configuration

func main() {
	var session *mgo.Session
	var err error

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if dbURL, ok := viper.Get("database.url").(string); ok {
		if session, err = mgo.Dial(dbURL); err != nil {
			panic(err)
		}
	}

	if dbName, ok := viper.Get("database.name").(string); ok {
		models.DB = session.DB(dbName)
	} else {
		panic("Invalid database name (wrong type)")
	}

	router := mux.NewRouter()

	router.HandleFunc("/hero", routes.GetHeroes).Methods("GET")
	router.HandleFunc("/hero", routes.AddHero).Methods("POST")
	router.HandleFunc("/hero", routes.ModifyHero).Methods("PUT")
	router.HandleFunc("/hero/{id}", routes.GetHero).Methods("GET")
	router.HandleFunc("/hero/{id}", routes.DeleteHero).Methods("DELETE")

	log.Printf("Running on port %v", viper.Get("port"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
