package services

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/junanda/golang-aa/models"
	"github.com/junanda/golang-aa/repository"
	"github.com/junanda/golang-aa/utils"
	"gorm.io/gorm"
)

type UserService interface {
	LoginUser(c *gin.Context, user models.User) (string, error)
	SignUp(c *gin.Context, user models.User) error
	LogOut(c *gin.Context, token string) error
}

type userServiceImpl struct {
	userRepo repository.UserRepository
	authRepo repository.AuthRepository
	jwtKey   []byte
}

func InitUserService(userRepo repository.UserRepository, ar repository.AuthRepository) UserService {
	jkey := []byte("my_secret_key")
	return &userServiceImpl{
		userRepo: userRepo,
		authRepo: ar,
		jwtKey:   jkey,
	}
}

func (u *userServiceImpl) LoginUser(c *gin.Context, user models.User) (string, error) {
	var (
		userExist models.User
		dataCache models.AuthData
		err       error
	)

	userExist, err = u.userRepo.FindByEmail(user.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not registered")
		}
		return "", err
	}

	errHash := utils.CompatreHashPassword(user.Password, userExist.Password)
	if !errHash {
		return "", errors.New("invalid password")
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	clain := &models.Claims{
		Role: userExist.Role,
		Uid:  userExist.Uid,
		StandardClaims: jwt.StandardClaims{
			Subject:   userExist.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clain)
	tokenString, err := token.SignedString(u.jwtKey)

	if err != nil {
		return "", errors.New("could not generate token")
	}

	dataCache.IdToken = userExist.Uid
	dataCache.JwtToken = tokenString
	dataCache.Role = userExist.Role

	err = u.authRepo.Save(c, dataCache, 0)
	if err != nil {
		log.Println("failed cache data")
	}

	// c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

	return tokenString, nil
}

func (u *userServiceImpl) SignUp(c *gin.Context, user models.User) error {
	var (
		exisingUser models.User
		errHash     error
	)

	exisingUser, err := u.userRepo.FindByEmail(user.Email)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
	}

	if exisingUser.ID != 0 {
		return errors.New("user already exists")
	}

	user.Password, errHash = utils.GenerateHashPassword(user.Password)
	if errHash != nil {
		return errors.New("could not generate password hash")
	}

	user.Uid = utils.GenerateUUID()

	err = u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userServiceImpl) LogOut(c *gin.Context, token string) error {
	// c.SetCookie("token", "", -1, "/", "localhost", false, true)
	claims, _ := utils.ParseToken(token)
	err := u.authRepo.DeleteKey(c, claims.Uid)
	if err != nil {
		log.Println("Cannot Remove Key: ", claims.Uid)
	}
	return nil
}
