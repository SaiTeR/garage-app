package database

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var RedisClient *redis.Client

// Инициализация Redis-клиента
func InitRedis() {
	// Создание клиента Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Адрес Redis-контейнера
		Password: "",           // Нет пароля (если не установлен)
		DB:       0,            // Используем базу данных с индексом 0
	})

	// Проверка подключения
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		panic("Failed to connect to Redis")
	}
	fmt.Println("Successfully connected to Redis")
}

// Функция для получения Redis клиента
func GetRedisClient() *redis.Client {
	return RedisClient
}
