package handler

import (
	"github.com/timac11/goph-keeper/internal/server/auth"
	"github.com/timac11/goph-keeper/internal/server/service"
)

type Handler struct {
	service    *service.Service
	jwtControl *auth.JWTControl
}

func NewHandler(service *service.Service, jwtControl *auth.JWTControl) *Handler {
	return &Handler{
		service:    service,
		jwtControl: jwtControl,
	}
}
