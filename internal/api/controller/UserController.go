package controller

import (
	"app/internal/api/types"
	"app/internal/service"
	"app/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.Service
}

func NewUserController(service service.Service) *UserController {

	return &UserController{
		service: service,
	}
}

func (u *UserController) Create(ctx *gin.Context) {
	userRequest := &types.CreateUserRequest{}
	// Bind the JSON body to the user struct
	err := ctx.BindJSON(&userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	util.PrintJSON(userRequest)

	err = u.service.UserService.Create(userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	userResponse := &types.CreateUserResponse{
		Email: userRequest.Email,
	}

	ctx.JSON(http.StatusOK, userResponse)
}

// func (u *UserController) FindByPk(ctx *gin.Context) {

// 	pk, err := strconv.Atoi(ctx.Param("pk"))
// 	if err != nil {
// 		fmt.Println("Error converting string to int int FindByPk controller:", err)
// 		return
// 	}

// 	user, err := u.userService.FindByPk(pk)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, user)
// }

// func (u *UserController) Update(ctx *gin.Context) {
// 	userRequest := &types.UpdateUserRequest{}

// 	err := ctx.BindJSON(&userRequest)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err = u.userService.Update(userRequest)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, map[string]string{"msg": "update ok"})
// }

// func (u *UserController) Delete(ctx *gin.Context) {
// 	userRequest := &types.DeleteUserRequest{}

// 	err := ctx.BindJSON(&userRequest)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err = u.userService.Delete(userRequest)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, map[string]string{"msg": "delete ok"})
// }
