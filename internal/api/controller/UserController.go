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
	err := ctx.BindJSON(&userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	util.PrintJSON(userRequest)

	err = u.userService.Create(userRequest)
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

func (u *UserController) Update(ctx *gin.Context) {
	userRequest := &dto.UpdateUserRequest{}

	err := ctx.BindJSON(&userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = u.userService.Update(userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{"msg": "update ok"})
}

func (u *UserController) Delete(ctx *gin.Context) {
	userRequest := &dto.DeleteUserRequest{}

	err := ctx.BindJSON(&userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = u.userService.Delete(userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{"msg": "delete ok"})
}
