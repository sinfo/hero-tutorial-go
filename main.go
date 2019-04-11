package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sinfo/go-tutorial/server"
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
