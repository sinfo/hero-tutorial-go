package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sinfo/go-tutorial/models"

	"github.com/gorilla/mux"
)

func GetHeroes(w http.ResponseWriter, r *http.Request) {
	heroes, err := models.GetHeroes()

	if err != nil {
		http.Error(w, "Unable to make query do database", http.StatusExpectationFailed)
	}

	json.NewEncoder(w).Encode(heroes)
}

func AddHero(w http.ResponseWriter, r *http.Request) {
	var hero models.Hero

	if err := json.NewDecoder(r.Body).Decode(&hero); err != nil {
		http.Error(w, "Invalid hero", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := hero.CreateHero(); err != nil {
		http.Error(w, "Could not create hero", http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(hero)
}

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

func ModifyHero(w http.ResponseWriter, r *http.Request) {
	var hero models.Hero

	if err := json.NewDecoder(r.Body).Decode(&hero); err != nil {
		http.Error(w, "Invalid hero", http.StatusBadRequest)
		return
	}

	newHero, err := models.ModifyHero(hero)

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

func DeleteHero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err = models.RemoveHero(id); err != nil {
		http.Error(w, "Failed to remove hero", http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusOK)
}
