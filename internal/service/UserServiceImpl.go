package service

import (
	"app/internal/model/dao"
	"app/internal/model/dto"
	"fmt"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserServiceImpl(db *gorm.DB) UserService {
	return &UserServiceImpl{
		db: db,
	}
}

func (u *UserServiceImpl) Create(createUserRequest *dto.CreateUserRequest) error {
	user := &dao.User{
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	u.db.Save(user)

	return nil
}

func (u *UserServiceImpl) FindByPk(pk int) (*dao.User, error) {
	user := dao.User{}

	result := u.db.First(&user, pk)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(user)

	return &user, nil
}

// func (u *UserServiceImpl) Update(user model.User) error

// func (u *UserServiceImpl) Delete(user model.User) error
