package service

import (
	"app/internal/model/dao"
	"app/internal/model/dto"
)

type UserService interface {
	Create(createUserRequest *dto.CreateUserRequest) error
	// Update(user model.User) error
	FindByPk(pk int) (*dao.User, error)
	// Delete(user model.User) error
}
