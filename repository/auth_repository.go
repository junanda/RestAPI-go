package repository

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/junanda/golang-aa/models"
	"github.com/redis/go-redis/v9"
)

type AuthRepository interface {
	Save(ctx *gin.Context, data models.AuthData, dataTtl int64) error
	GetData(ctx *gin.Context, dataKey string) (models.AuthData, error)
	DeleteKey(ctx *gin.Context, dataKey string) error
}

type authRepository struct {
	client *redis.Client
}

func NewAuthRepository(cl *redis.Client) AuthRepository {
	return &authRepository{
		client: cl,
	}
}

func (a *authRepository) Save(ctx *gin.Context, data models.AuthData, dataTtl int64) error {
	redisData, _ := json.Marshal(data)
	err := a.client.Set(ctx, data.IdToken, redisData, 0).Err()
	if err != nil {
		log.Println("error set data on redis: ", err.Error())
	}

	return err
}

func (a *authRepository) GetData(ctx *gin.Context, dataKey string) (models.AuthData, error) {
	var (
		dataAuth models.AuthData
	)
	// dataR := a.client.HGetAll(ctx, dataKey).Val()
	dataR, err := a.client.Get(ctx, dataKey).Result()
	if err != nil && err != redis.Nil {
		return dataAuth, err
	}

	if err := json.Unmarshal([]byte(dataR), &dataAuth); err != nil {
		return dataAuth, err
	}

	return dataAuth, nil
}

func (a *authRepository) DeleteKey(ctx *gin.Context, datakey string) error {
	_, err := a.client.Del(ctx, datakey).Result()
	if err != nil {
		return err
	}

	return nil
}
