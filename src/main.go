// Package classification Petstore API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost:8000
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
// go:generate swagger generate spec
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sinfo/hero-tutorial-go/src/server"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("GO_TUTORIAL")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	dbURL := viper.GetString("DB_URL")
	dbName := viper.GetString("DB_NAME")

	server.InitServer(dbURL, dbName)

	log.Printf("Running on port %v", viper.Get("port"))
	log.Fatal(http.ListenAndServe(":8000", server.ServerInstance))
}
