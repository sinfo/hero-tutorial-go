package routes

import (
	"net/http"
)

// GetSwagger is the handler that gets all the heroes from the database
func GetSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./swagger.yaml")
}
