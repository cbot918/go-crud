package main

import (
	"app/database"
	"app/internal/api"
	"app/internal/api/controller"
	"app/internal/configs"
	"app/internal/repository"
	"log"

	"app/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// init config
	config, err := configs.New()
	if err != nil {
		log.Fatal(err)
	}

	// init database
	database, err := database.NewMysql(
		config.Connections.MySQL.User,
		config.Connections.MySQL.Password,
		config.Connections.MySQL.Host,
		config.Connections.MySQL.DB,
	)
	if err != nil {
		log.Fatal(err)
	}

	// init controller / service
	repository := repository.NewRepository(database)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	// init server
	server := api.NewHTTPServer(gin.Default(), controller)

	// run server
	err = server.Router.Run(":" + config.Server.Port)
	if err != nil {
		log.Fatal(err)
	}
}
