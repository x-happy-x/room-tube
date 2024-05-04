package main

import (
	"log"
	"tube/pkg/api"
	"tube/pkg/config"
	"tube/pkg/middleware"
	"tube/pkg/repository/postgre"
	"tube/pkg/service"
)

func main() {
	appConfig, err := config.LoadApplicationConfig("./application.yaml")
	if err != nil {
		log.Fatal(err)
	}

	repository, err := postgre.NewRepository(appConfig.Database)
	if err != nil {
		log.Fatal(err)
	}

	authService := service.NewAuthService(repository, appConfig)
	userService := service.NewUserService(repository, appConfig)
	services := service.Service{
		Auth: authService,
		User: userService,
	}

	middleware := middleware.NewMiddleware(appConfig)

	roomTube := api.NewAPI(appConfig)
	log.Fatal(roomTube.Run())
}
