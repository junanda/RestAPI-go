package database

import (
	"fmt"
	"log"

	"github.com/junanda/golang-aa/config"
	"github.com/redis/go-redis/v9"
)

type RedisDB interface {
	Database
	GetDb() *redis.Client
}

type RedisDatabase struct {
	cfg    config.Config
	client *redis.Client
}

func NewRedisDB(cfg config.Config) RedisDB {
	return &RedisDatabase{
		cfg: cfg,
	}
}

func (r *RedisDatabase) Connect() {
	dns := fmt.Sprintf("%s:%s", r.cfg.HostRedis, r.cfg.PortRedis)
	r.client = redis.NewClient(&redis.Options{
		Addr:     dns,
		Password: r.cfg.PasswordRedis,
		DB:       0,
	})
	log.Println("Redis connection success...")
}

func (r *RedisDatabase) Close() {
	r.client.Close()
	log.Println("Redis connection close...")
}

func (r *RedisDatabase) GetDb() *redis.Client {
	return r.client
}
