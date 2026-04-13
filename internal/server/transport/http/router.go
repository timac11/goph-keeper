package http

import (
	"github.com/gorilla/mux"
	"github.com/timac11/goph-keeper/internal/server/auth"
	"github.com/timac11/goph-keeper/internal/server/transport/http/handler"
	"github.com/timac11/goph-keeper/internal/server/transport/http/middleware"
)

func Init(h handler.Handler, jwtControl *auth.JWTControl) *mux.Router {
	r := mux.NewRouter()

	m := middleware.NewMiddleware(middleware.Params{JwtControl: jwtControl})

	r.Use(m.RequestLoggerMiddleware)

	privateRouter := r.PathPrefix("/api").Subrouter()
	privateRouter.Use(m.AuthCheckMiddleware)

	publicRouter := r.PathPrefix("/api").Subrouter()
	publicRouter.Use(m.NotAuthCheckMiddleware)
	publicRouter.HandleFunc("/login", h.Login).Methods("POST")
	publicRouter.HandleFunc("/register", h.Login).Methods("POST")

	return r
}
