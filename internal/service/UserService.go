package service

import (
	"app/internal/model/dao"
	"app/internal/model/dto"
)

type UserService interface {
	Create(createUserRequest *dto.CreateUserRequest) error
	Update(user *dto.UpdateUserRequest) error
	FindByPk(pk int) (*dao.User, error)
	Delete(user *dto.DeleteUserRequest) error
}
