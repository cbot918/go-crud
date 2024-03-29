package controller

import (
	"app/internal/model/dto"
	"app/internal/service"
	"app/internal/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {

	return &UserController{
		userService: userService,
	}
}

func (u *UserController) Create(ctx *gin.Context) {
	userRequest := &dto.CreateUserRequest{}
	// Bind the JSON body to the user struct
	if err := ctx.BindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	util.PrintJSON(userRequest)

	err := u.userService.Create(userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	userResponse := &dto.CreateUserResponse{
		Email: userRequest.Email,
	}

	ctx.JSON(http.StatusOK, userResponse)
}

func (u *UserController) FindByPk(ctx *gin.Context) {

	pk, err := strconv.Atoi(ctx.Param("pk"))
	if err != nil {
		fmt.Println("Error converting string to int int FindByPk controller:", err)
		return
	}

	user, err := u.userService.FindByPk(pk)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// func (u *UserController) Update(user model.User) error

// func (u *UserController) Delete(user model.User) error
