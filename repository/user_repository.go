package repository

import (
	"github.com/junanda/golang-aa/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (models.User, error)
	CreateUser(data models.User) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (u *userRepositoryImpl) FindByEmail(email string) (models.User, error) {
	var (
		existingUser models.User
	)

	err := u.db.Where("email = ?", email).First(&existingUser).Error
	if err != nil {
		return existingUser, err
	}

	return existingUser, nil
}

func (u *userRepositoryImpl) CreateUser(data models.User) error {
	err := u.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}
