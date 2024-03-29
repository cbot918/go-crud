package main

import (
	"app/database"
	"app/internal/api"
	"app/internal/api/controller"
	"app/internal/configs"
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
	userService := service.NewUserServiceImpl(database)
	userController := controller.NewUserController(userService)

	// init server
	server := api.NewHTTPServer(gin.Default(), userController)

	// run server
	err = server.Router.Run(":" + config.Server.Port)
	if err != nil {
		log.Fatal(err)
	}
}
