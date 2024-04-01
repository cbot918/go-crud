package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository UserRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		UserRepository: NewUserRepositoryImpl(db),
	}
}
