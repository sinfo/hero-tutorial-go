package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sinfo/hero-tutorial-go/src/models"

	"github.com/gorilla/mux"
)

// GetHeroes is the handler that gets all the heroes from the database
// swagger:route GET /hero heroes getHeroes
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: Hero
//       422
func GetHeroes(w http.ResponseWriter, r *http.Request) {
	heroes, err := models.GetHeroes()

	if err != nil {
		http.Error(w, "Unable to make query do database", http.StatusExpectationFailed)
	}

	json.NewEncoder(w).Encode(heroes)
}

func AddHero(w http.ResponseWriter, r *http.Request) {
	var hero *models.Hero
	var err error

	if hero, err = models.HeroFromBody(r.Body); err != nil {
		http.Error(w, "Invalid hero", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err = hero.CreateHero(); err != nil {
		http.Error(w, "Could not create hero", http.StatusConflict)
		return
	}

	json.NewEncoder(w).Encode(hero)
}

// GetHero is the handler that gets a specific hero from the database
func GetHero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	hero, err := models.GetHero(id)

	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	if hero != nil {
		json.NewEncoder(w).Encode(hero)
		return
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

// ModifyHero is the handler that renames a hero from the database
func ModifyHero(w http.ResponseWriter, r *http.Request) {
	var hero *models.Hero
	var err error

	if hero, err = models.HeroFromBody(r.Body); err != nil {
		http.Error(w, "Invalid hero", http.StatusBadRequest)
		return
	}

	newHero, err := models.ModifyHero(*hero)

	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	if newHero != nil {
		json.NewEncoder(w).Encode(newHero)
		return
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

// DeleteHero is the handler that removes a hero from the database
func DeleteHero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err = models.RemoveHero(id); err != nil {
		http.Error(w, "Failed to remove hero", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
