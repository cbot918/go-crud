package api

import (
	"app/internal/api/controller"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Router         *gin.Engine
	userController *controller.UserController
}

func NewHTTPServer(router *gin.Engine, userController *controller.UserController) *HTTPServer {

	user := router.Group("/v1/user")
	{
		user.POST("/create", userController.Create)
		user.GET("/find-by-pk/:pk", userController.FindByPk)
		user.POST("/update", userController.Update)
		user.POST("/delete", userController.Delete)
	}

	return &HTTPServer{
		Router:         router,
		userController: userController,
	}

}
