package routes_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/azbshiri/common/test"
	"github.com/sinfo/go-tutorial/models"
	"github.com/sinfo/go-tutorial/server"

	"github.com/spf13/viper"
	"gotest.tools/assert"
)

func TestMain(m *testing.M) {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	dbURL := viper.GetString("DB_TEST_URL")
	dbName := viper.GetString("DB_TEST_NAME")

	server.InitServer(dbURL, dbName)

	retCode := m.Run()

	server.ServerInstance.DB.DropDatabase()

	os.Exit(retCode)
}

func TestAddHero(t *testing.T) {
	hero := &models.Hero{ID: 1, Name: "my_test_name"}
	b, err := json.Marshal(hero)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := test.DoRequest(server.ServerInstance, "POST", "/hero", bytes.NewBuffer(b))
	assert.NilError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)
}
