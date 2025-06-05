package repository

import (
	"github.com/sharipov/sunnatillo/academy-backend/internal/models"
	"github.com/sharipov/sunnatillo/academy-backend/pkg/dto"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository UserRepository) Create(user *models.User) (uint, error) {
	if err := repository.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (repository UserRepository) Get(id uint) (*models.User, error) {
	var user models.User
	if err := repository.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository UserRepository) Update(user *models.User) error {
	if err := repository.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repository UserRepository) Delete(id uint) error {
	if err := repository.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repository UserRepository) GetAll(filter dto.UserFilter) ([]*models.User, error) {
	var users []*models.User
	if err := repository.db.Where("first_name LIKE ?", "%"+filter.Search+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
