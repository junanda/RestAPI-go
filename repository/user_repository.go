package repository

import (
	"errors"

	"github.com/junanda/golang-aa/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (models.User, error)
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

	u.db.Where("email = ?", email).First(&existingUser)
	if existingUser.ID == 0 {
		return existingUser, errors.New("user does not exist")
	}

	return existingUser, nil
}
