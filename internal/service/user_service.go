package service

import (
	"app/internal/api/types"
	"app/internal/model"
	"app/internal/repository"
)

type UserService interface {
	Create(createUserRequest *types.CreateUserRequest) error
	// Update(user *types.UpdateUserRequest) error
	// FindByPk(pk int) (*model.User, error)
	// Delete(user *types.DeleteUserRequest) error
}

type UserServiceImpl struct {
	repo repository.Repository
}

func NewUserServiceImpl(repository repository.Repository) UserService {
	return &UserServiceImpl{
		repo: repository,
	}
}

func (u *UserServiceImpl) Create(createUserRequest *types.CreateUserRequest) error {

	user := &model.User{
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	err := u.repo.UserRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

// func (u *UserServiceImpl) Update(user *types.UpdateUserRequest) error {

// 	u.db.
// 		Model(&model.User{}).
// 		Where("id=?", user.Id).
// 		Updates(model.User{Password: user.Password})

// 	return nil
// }

// func (u *UserServiceImpl) FindByPk(pk int) (*model.User, error) {
// 	user := model.User{}

// 	result := u.db.First(&user, pk)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	fmt.Println(user)

// 	return &user, nil
// }

// func (u *UserServiceImpl) Delete(user *types.DeleteUserRequest) error {

// 	u.db.Delete(&model.User{}, user.Id)

// 	return nil
// }
