package http

import (
	"github.com/gorilla/mux"
	"github.com/timac11/goph-keeper/internal/server/auth"
	"github.com/timac11/goph-keeper/internal/server/transport/http/handler"
	"github.com/timac11/goph-keeper/internal/server/transport/http/middleware"
)

// @title Gophkeeper server API
// @version 1.0
// @description This is a Gophkeeper swagger API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /
func Init(h handler.Handler, jwtControl *auth.JWTControl) *mux.Router {
	r := mux.NewRouter()

	m := middleware.NewMiddleware(middleware.Params{JwtControl: jwtControl})

	r.Use(m.RequestLoggerMiddleware)

	privateRouter := r.PathPrefix("/api").Subrouter()
	privateRouter.Use(m.AuthCheckMiddleware)
	privateRouter.HandleFunc("/secrets", h.GetSecretsList).Methods("GET")
	privateRouter.HandleFunc("/secrets/{id}", h.GetSecret).Methods("GET")
	privateRouter.HandleFunc("/secrets/{id}", h.DeleteSecret).Methods("DELETE")
	privateRouter.HandleFunc("/secrets", h.CreateSecret).Methods("POST")

	publicRouter := r.PathPrefix("/api").Subrouter()
	publicRouter.Use(m.NotAuthCheckMiddleware)
	publicRouter.HandleFunc("/login", h.Login).Methods("POST")
	publicRouter.HandleFunc("/register", h.Register).Methods("POST")

	return r
}
