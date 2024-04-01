package api

import (
	"app/internal/api/controller"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Router     *gin.Engine
	controller controller.Controller
}

func NewHTTPServer(router *gin.Engine, controller controller.Controller) *HTTPServer {

	user := router.Group("/v1/user")
	{
		user.POST("/create", controller.UserController.Create)
		// user.GET("/find-by-pk/:pk", userController.FindByPk)
		// user.POST("/update", userController.Update)
		// user.POST("/delete", userController.Delete)
	}

	return &HTTPServer{
		Router:     router,
		controller: controller,
	}

}
