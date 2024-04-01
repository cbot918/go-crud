package service

import "app/internal/repository"

type Service struct {
	UserService UserService
	// postService PostService
}

func NewService(repository repository.Repository) Service {
	return Service{
		UserService: NewUserServiceImpl(repository),
	}
}
