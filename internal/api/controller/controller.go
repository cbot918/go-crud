package controller

import "app/internal/service"

type Controller struct {
	UserController UserController
}

func NewController(service service.Service) Controller {
	return Controller{
		UserController: *NewUserController(service),
	}
}
