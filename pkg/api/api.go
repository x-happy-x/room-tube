package api

import (
	"net/http"
	"tube/pkg/config"
	"tube/pkg/middleware"
	"tube/pkg/service"
)

type API interface {
	Authorize(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)

	Run() error
}

type roomTubeAPI struct {
	config     *config.Api
	handler    http.Handler
	service    service.Service
	middleware middleware.Middleware
}

func NewAPI(appConfig *config.Application, handlers ...interface{}) API {
	api := roomTubeAPI{
		config:     &appConfig.Api,
		middleware: middleware.NewMiddleware(appConfig),
	}
	api.handler = NewRouter(&api)
	return &api
}

func (api *roomTubeAPI) Run() error {
	return http.ListenAndServe(api.config.Host+":"+api.config.Port, api.handler)
}
