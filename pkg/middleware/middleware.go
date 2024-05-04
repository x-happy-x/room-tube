package middleware

import (
	"net/http"
	"tube/pkg/config"
	"tube/pkg/service"
)

type Middleware interface {
	AuthMiddleware(next http.Handler) http.Handler
}

type Struct struct {
	config   *config.Application
	services service.Service
}

func NewMiddleware(appConfig *config.Application, services service.Service) Middleware {
	return &Struct{
		config:   appConfig,
		services: services,
	}
}
