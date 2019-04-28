package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sinfo/hero-tutorial-go/src/models"

	"github.com/gorilla/mux"
)

// GetHeroes is the handler that gets all the heroes from the database
func GetHeroes(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /hero heroes GetHeroes
	//
	// Returns all heroes
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: Gets all heroes
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/Hero"

	heroes, err := models.GetHeroes()

	if err != nil {
		http.Error(w, "Unable to make query do database", http.StatusExpectationFailed)
	}

	json.NewEncoder(w).Encode(heroes)
}

// AddHero is the handler that adds a hero
func AddHero(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /hero heroes AddHero
	//
	// Returns the created hero
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: hero
	//   in: body
	//   description: Hero body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Hero"
	// responses:
	//   '200':
	//    description: Created hero
	//    schema:
	//      "$ref": "#/definitions/Hero"

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
	// swagger:operation GET /hero/{id} heroes GetHero
	//
	// Returns a specific hero
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - in: path
	//   name: id
	//   description: id of hero
	//   required: true
	//   type: integer
	// responses:
	//   '200':
	//     description: Gets a specific hero
	//     schema:
	//       "$ref": "#/definitions/Hero"

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
	// swagger:operation PUT /hero heroes ModifyHero
	//
	// Returns the modified hero
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: hero
	//   in: body
	//   description: Hero body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Hero"
	// responses:
	//   '200':
	//    description: Modified hero
	//    schema:
	//      "$ref": "#/definitions/Hero"

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
	// swagger:operation DELETE /hero/{id} heroes DeleteHero
	//
	// Deletes a hero
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - in: path
	//   name: id
	//   description: id of hero
	//   required: true
	//   type: integer
	// responses:
	//   '200':
	//     description: Returns the deleted hero
	//     schema:
	//       "$ref": "#/definitions/Hero"

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
