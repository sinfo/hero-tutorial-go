package routes_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/sinfo/hero-tutorial-go/src/models"
	"github.com/sinfo/hero-tutorial-go/src/server"

	"github.com/spf13/viper"
	"gotest.tools/assert"
)

func executeRequest(method string, path string, payload io.Reader) (*httptest.ResponseRecorder, error) {
	req, errReq := http.NewRequest(method, path, payload)

	if errReq != nil {
		return nil, errReq
	}

	rr := httptest.NewRecorder()
	server.ServerInstance.ServeHTTP(rr, req)

	return rr, nil
}

func TestMain(m *testing.M) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("GO_TUTORIAL")
	viper.AutomaticEnv()

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
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	hero := &models.Hero{ID: 1, Name: "my_test_name"}
	b, errMarshal := json.Marshal(hero)

	assert.NilError(t, errMarshal)

	res, err := executeRequest("POST", "/hero", bytes.NewBuffer(b))
	assert.NilError(t, err)

	assert.Equal(t, res.Code, http.StatusOK)
}

func TestAddHeroDuplicate(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	hero := &models.Hero{ID: 1, Name: "my_test_name"}
	duplicated := &models.Hero{ID: 1, Name: "my_other_test_name"}
	heroB, errMarshal1 := json.Marshal(hero)
	duplicatedB, errMarshal2 := json.Marshal(duplicated)

	assert.NilError(t, errMarshal1)
	assert.NilError(t, errMarshal2)

	res1, err1 := executeRequest("POST", "/hero", bytes.NewBuffer(heroB))
	res2, _ := executeRequest("POST", "/hero", bytes.NewBuffer(duplicatedB))

	assert.NilError(t, err1)
	assert.Equal(t, res1.Code, http.StatusOK)

	assert.Equal(t, res2.Code, http.StatusConflict)
}

func TestAddHeroInvalidType(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	type WrongHero struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	hero := &WrongHero{ID: "1", Name: "my_wrong_type_id"}
	heroB, errMarshal := json.Marshal(hero)

	assert.NilError(t, errMarshal)

	res, _ := executeRequest("POST", "/hero", bytes.NewBuffer(heroB))

	assert.Equal(t, res.Code, http.StatusBadRequest)
}
func TestAddHeroInvalidFormat(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	type WrongHero struct {
		ID       int    `json:"id"`
		FakeName string `json:"fakeName"`
	}

	hero := &WrongHero{ID: 1, FakeName: "my_wrong_format"}
	heroB, errMarshal := json.Marshal(hero)

	assert.NilError(t, errMarshal)

	res, _ := executeRequest("POST", "/hero", bytes.NewBuffer(heroB))

	assert.Equal(t, res.Code, http.StatusBadRequest)
}

func TestGetsHeroes(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	var heroes []models.Hero
	hero1 := &models.Hero{ID: 1, Name: "some hero"}
	hero2 := &models.Hero{ID: 2, Name: "some other hero"}

	errCreateHero1 := hero1.CreateHero()
	assert.NilError(t, errCreateHero1)

	errCreateHero2 := hero2.CreateHero()
	assert.NilError(t, errCreateHero2)

	res, err := executeRequest("GET", "/hero", nil)
	assert.NilError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)

	json.NewDecoder(res.Body).Decode(&heroes)

	assert.Equal(t, hero1.IsIn(heroes), true)
	assert.Equal(t, hero2.IsIn(heroes), true)

}
func TestGetHero(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	createdHero := &models.Hero{ID: 1, Name: "some hero"}
	var queriedHero models.Hero

	errCreateHero := createdHero.CreateHero()
	assert.NilError(t, errCreateHero)

	res, err := executeRequest("GET", fmt.Sprintf("/hero/%v", createdHero.ID), nil)
	assert.NilError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)

	json.NewDecoder(res.Body).Decode(&queriedHero)

	assert.Equal(t, createdHero.Equals(queriedHero), true)
}

func TestGetHeroInvalidID(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()
	res, _ := executeRequest("GET", fmt.Sprintf("/hero/not_a_number"), nil)
	assert.Equal(t, res.Code, http.StatusBadRequest)
}

func TestGetHeroNotFound(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	createdHero := &models.Hero{ID: 1, Name: "some hero"}

	res, _ := executeRequest("GET", fmt.Sprintf("/hero/%v", createdHero.ID), nil)
	assert.Equal(t, res.Code, http.StatusNotFound)
}

func TestDeleteHero(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	var heroes []models.Hero
	hero1 := &models.Hero{ID: 1, Name: "some hero"}
	hero2 := &models.Hero{ID: 2, Name: "some other hero"}

	errCreateHero1 := hero1.CreateHero()
	assert.NilError(t, errCreateHero1)

	errCreateHero2 := hero2.CreateHero()
	assert.NilError(t, errCreateHero2)

	res, err := executeRequest("DELETE", fmt.Sprintf("/hero/%v", hero1.ID), nil)
	assert.NilError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)

	res, err = executeRequest("GET", "/hero", nil)

	json.NewDecoder(res.Body).Decode(&heroes)

	assert.Equal(t, hero1.IsIn(heroes), false)
	assert.Equal(t, hero2.IsIn(heroes), true)
}

func TestDeleteHeroInvalidID(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	var heroes []models.Hero
	hero1 := &models.Hero{ID: 1, Name: "some hero"}
	hero2 := &models.Hero{ID: 2, Name: "some other hero"}

	errCreateHero1 := hero1.CreateHero()
	assert.NilError(t, errCreateHero1)

	errCreateHero2 := hero2.CreateHero()
	assert.NilError(t, errCreateHero2)

	res, _ := executeRequest("DELETE", fmt.Sprintf("/hero/invalid_id"), nil)
	assert.Equal(t, res.Code, http.StatusBadRequest)

	res, _ = executeRequest("GET", "/hero", nil)

	json.NewDecoder(res.Body).Decode(&heroes)

	assert.Equal(t, hero1.IsIn(heroes), true)
	assert.Equal(t, hero2.IsIn(heroes), true)
}

func TestDeleteHeroNotFound(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	var heroes []models.Hero
	hero1 := &models.Hero{ID: 1, Name: "some hero"}
	hero2 := &models.Hero{ID: 2, Name: "some other hero"}

	errCreateHero1 := hero1.CreateHero()
	assert.NilError(t, errCreateHero1)

	errCreateHero2 := hero2.CreateHero()
	assert.NilError(t, errCreateHero2)

	res, _ := executeRequest("DELETE", fmt.Sprintf("/hero/%v", hero1.ID+hero2.ID), nil)
	assert.Equal(t, res.Code, http.StatusNotFound)

	res, _ = executeRequest("GET", "/hero", nil)

	json.NewDecoder(res.Body).Decode(&heroes)

	assert.Equal(t, hero1.IsIn(heroes), true)
	assert.Equal(t, hero2.IsIn(heroes), true)
}

func TestModifyHero(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	hero := &models.Hero{ID: 1, Name: "my_test_name"}
	var queriedHero models.Hero

	errHero := hero.CreateHero()
	assert.NilError(t, errHero)

	hero.Name = "changed name"
	b, errMarshal := json.Marshal(hero)

	assert.NilError(t, errMarshal)

	res, err := executeRequest("PUT", "/hero", bytes.NewBuffer(b))
	json.NewDecoder(res.Body).Decode(&queriedHero)

	assert.NilError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, hero.Equals(queriedHero), true)
}

func TestModifyHeroInvalidFormat(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	hero := &models.Hero{ID: 1, Name: "my_test_name"}

	errHero := hero.CreateHero()
	assert.NilError(t, errHero)

	type WrongHero struct {
		ID       int    `json:"id"`
		FakeName string `json:"fakeName"`
	}

	wrongHero := &WrongHero{ID: 1, FakeName: "my_wrong_format"}
	heroB, errMarshal := json.Marshal(wrongHero)

	assert.NilError(t, errMarshal)

	res, _ := executeRequest("PUT", "/hero", bytes.NewBuffer(heroB))
	assert.Equal(t, res.Code, http.StatusBadRequest)
}

func TestModifyHeroNotFound(t *testing.T) {
	defer server.ServerInstance.DB.C("heroes").DropCollection()

	hero := &models.Hero{ID: 1, Name: "my_test_name"}

	errHero := hero.CreateHero()
	assert.NilError(t, errHero)

	hero.Name = "changed name"
	hero.ID += 1
	b, errMarshal := json.Marshal(hero)

	assert.NilError(t, errMarshal)

	res, _ := executeRequest("PUT", "/hero", bytes.NewBuffer(b))
	assert.Equal(t, res.Code, http.StatusNotFound)
}
