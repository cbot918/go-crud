package repository

import (
	"app/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	FindByPk(pk int) (*model.User, error)
	Delete(userser *model.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {

	return &UserRepositoryImpl{
		db: db,
	}
}
func (u *UserRepositoryImpl) Create(user *model.User) error {

	u.db.Save(user)

	return nil
}
func (u *UserRepositoryImpl) Update(user *model.User) error {
	u.db.
		Model(&model.User{}).
		Where("id=?", user.Id).
		Updates(model.User{Password: user.Password})

	return nil
}
func (u *UserRepositoryImpl) FindByPk(pk int) (*model.User, error) {
	user := model.User{}

	result := u.db.First(&user, pk)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
func (u *UserRepositoryImpl) Delete(user *model.User) error {

	u.db.Delete(&model.User{}, user.Id)

	return nil
}
