package service

import (
	"app/internal/api/types"
	"app/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type PostService interface {
	Create(createUserRequest *types.CreateUserRequest) error
	Update(user *types.UpdateUserRequest) error
	FindByPk(pk int) (*model.User, error)
	Delete(user *types.DeleteUserRequest) error
}

type PostServiceImpl struct {
	db *gorm.DB
}

func NewPostServiceImpl(db *gorm.DB) PostService {
	return &PostServiceImpl{
		db: db,
	}
}

func (u *PostServiceImpl) Create(createUserRequest *types.CreateUserRequest) error {
	user := &model.User{
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	u.db.Save(user)

	return nil
}

func (u *PostServiceImpl) Update(user *types.UpdateUserRequest) error {

	u.db.
		Model(&model.User{}).
		Where("id=?", user.Id).
		Updates(model.User{Password: user.Password})

	return nil
}

func (u *PostServiceImpl) FindByPk(pk int) (*model.User, error) {
	user := model.User{}

	result := u.db.First(&user, pk)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(user)

	return &user, nil
}

func (u *PostServiceImpl) Delete(user *types.DeleteUserRequest) error {

	u.db.Delete(&model.User{}, user.Id)

	return nil
}
