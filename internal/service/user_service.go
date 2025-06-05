package service

import (
	"github.com/sharipov/sunnatillo/academy-backend/internal/models"
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
	user := models.User{
		FirstName:  createDto.FirstName,
		LastName:   createDto.LastName,
		MiddleName: createDto.MiddleName,
		Email:      createDto.Email,
		Phone:      createDto.Phone,
	}
	service.userRepository.Create(&user)
	return user.ID, nil
}
