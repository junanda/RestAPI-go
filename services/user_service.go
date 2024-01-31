package services

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/junanda/golang-aa/models"
	"github.com/junanda/golang-aa/repository"
	"github.com/junanda/golang-aa/utils"
)

type UserService interface {
	LoginUser(c *gin.Context, user models.User) error
}

type userServiceImpl struct {
	userRepo repository.UserRepository
	jwtKey   []byte
}

func InitUserService(userRepo repository.UserRepository) UserService {
	jkey := []byte("my_secret_key")
	return &userServiceImpl{
		userRepo: userRepo,
		jwtKey:   jkey,
	}
}

func (u *userServiceImpl) LoginUser(c *gin.Context, user models.User) error {
	var (
		userExist models.User
		err       error
	)

	userExist, err = u.userRepo.FindByEmail(user.Email)
	if err != nil {
		return err
	}

	errHash := utils.CompatreHashPassword(user.Password, userExist.Password)
	if !errHash {
		return errors.New("invalid password")
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	clain := &models.Claims{
		Role: userExist.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   userExist.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clain)
	tokenString, err := token.SignedString(u.jwtKey)

	if err != nil {
		return errors.New("could not generate token")
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

	return nil
}
