package routes

import (
	"net/http"
)

// GetSwagger is the handler that gets all the heroes from the database
// swagger:route GET /swagger swagger_config GetSwagger
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
func GetSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./swagger.yaml")
}
