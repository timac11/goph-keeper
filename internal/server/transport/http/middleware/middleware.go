package middleware

import (
	"github.com/timac11/goph-keeper/internal/server/auth"
)

type Params struct {
	JwtControl *auth.JWTControl
}

type Middleware struct {
	jwtControl *auth.JWTControl
}

func NewMiddleware(params Params) *Middleware {
	mw := Middleware{
		jwtControl: params.JwtControl,
	}
	return &mw
}
