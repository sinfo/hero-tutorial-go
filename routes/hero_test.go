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
	server.ServerInstance.DB.DropDatabase()

	retCode := m.Run()

	os.Exit(retCode)
}

func TestAddHero(t *testing.T) {
	hero := &models.Hero{ID: 1, Name: "my_test_name"}
	b, errMarshal := json.Marshal(hero)

	assert.NilError(t, errMarshal)

	res, err := test.DoRequest(server.ServerInstance, "POST", "/hero", bytes.NewBuffer(b))
	assert.NilError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)

	server.ServerInstance.DB.C("heroes").DropCollection()
}

func TestGetsHeroes(t *testing.T) {
	var heroes []models.Hero
	hero1 := &models.Hero{ID: 1, Name: "some hero"}
	hero2 := &models.Hero{ID: 2, Name: "some other hero"}

	errCreateHero1 := hero1.CreateHero()
	assert.NilError(t, errCreateHero1)

	errCreateHero2 := hero2.CreateHero()
	assert.NilError(t, errCreateHero2)

	res, err := test.DoRequest(server.ServerInstance, "GET", "/hero", nil)
	assert.NilError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)

	json.NewDecoder(res.Body).Decode(&heroes)

	assert.Equal(t, hero1.IsIn(heroes), true)
	assert.Equal(t, hero2.IsIn(heroes), true)

	server.ServerInstance.DB.C("heroes").DropCollection()
}
func TestGetHero(t *testing.T) {
	createdHero := &models.Hero{ID: 1, Name: "some hero"}
	var queriedHero models.Hero

	errCreateHero := createdHero.CreateHero()
	assert.NilError(t, errCreateHero)

	res, err := test.DoRequest(server.ServerInstance, "GET", fmt.Sprintf("/hero/%v", createdHero.ID), nil)
	assert.NilError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)

	json.NewDecoder(res.Body).Decode(&queriedHero)

	assert.Equal(t, createdHero.Equals(queriedHero), true)

	server.ServerInstance.DB.C("heroes").DropCollection()
}

func TestDeleteHero(t *testing.T) {
	var heroes []models.Hero
	hero1 := &models.Hero{ID: 1, Name: "some hero"}
	hero2 := &models.Hero{ID: 2, Name: "some other hero"}

	errCreateHero1 := hero1.CreateHero()
	assert.NilError(t, errCreateHero1)

	errCreateHero2 := hero2.CreateHero()
	assert.NilError(t, errCreateHero2)

	res, err := test.DoRequest(server.ServerInstance, "DELETE", fmt.Sprintf("/hero/%v", hero1.ID), nil)
	assert.NilError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)

	res, err = test.DoRequest(server.ServerInstance, "GET", "/hero", nil)

	json.NewDecoder(res.Body).Decode(&heroes)

	assert.Equal(t, hero1.IsIn(heroes), false)
	assert.Equal(t, hero2.IsIn(heroes), true)

	server.ServerInstance.DB.C("heroes").DropCollection()
}

func TestModifyHero(t *testing.T) {
	hero := &models.Hero{ID: 1, Name: "my_test_name"}
	var queriedHero models.Hero

	errHero := hero.CreateHero()
	assert.NilError(t, errHero)

	hero.Name = "changed name"
	b, errMarshal := json.Marshal(hero)

	assert.NilError(t, errMarshal)

	res, err := test.DoRequest(server.ServerInstance, "PUT", "/hero", bytes.NewBuffer(b))
	json.NewDecoder(res.Body).Decode(&queriedHero)

	assert.NilError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, hero.Equals(queriedHero), true)

	server.ServerInstance.DB.C("heroes").DropCollection()
}
