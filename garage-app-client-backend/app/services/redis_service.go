package services

import (
	"context"
	"garage-app-client-backend/database"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisService struct {
	Client *redis.Client
}

// Создание нового RedisService
func NewRedisService() *RedisService {
	return &RedisService{Client: database.GetRedisClient()}
}

// Сохранение токена в Redis
func (s *RedisService) StoreToken(userID string, token string, expiration time.Duration) error {
	ctx := context.Background()
	return s.Client.Set(ctx, token, userID, expiration).Err()
}

//
//// Получение токена из Redis
//func (s *RedisService) GetToken(userID string) (string, error) {
//	ctx := context.Background()
//	return s.Client.Get(ctx, userID).Result()
//}

func (s *RedisService) CheckToken(ctx *gin.Context, token string) (int64, error) {
	return s.Client.Exists(ctx, token).Result()
}

// Удаление токена из Redis
func (s *RedisService) DeleteToken(userID string) error {
	ctx := context.Background()
	return s.Client.Del(ctx, userID).Err()
}
