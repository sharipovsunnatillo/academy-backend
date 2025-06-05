package service

import (
	"github.com/sharipov/sunnatillo/academy-backend/internal/repository"
	"github.com/sharipov/sunnatillo/academy-backend/pkg/dto"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (service UserService) Register(createDto dto.UserCreateDto) (uint, error) {
	return 0, nil
}
