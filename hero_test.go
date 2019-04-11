package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/azbshiri/common/test"
	"github.com/sinfo/go-tutorial/models"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var testServer *server

func TestMain(m *testing.M) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	dbURL := viper.GetString("DB_TEST_URL")
	dbName := viper.GetString("DB_TEST_NAME")

	testServer = newServer(dbURL, dbName)

	retCode := m.Run()

	testServer.db.DropDatabase()

	os.Exit(retCode)
}

func TestAddHero(t *testing.T) {
	hero := &models.Hero{ID: 1, Name: "my_test_name"}
	b, err := json.Marshal(hero)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := test.DoRequest(testServer, "POST", "/hero", bytes.NewBuffer(b))
	assert.NoError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)
}
