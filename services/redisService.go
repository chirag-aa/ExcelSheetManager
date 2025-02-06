package services

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"excelsheetmanager.com/models"
	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	redisClient *redis.Client
}

var ctx = context.Background()

func NewRedisService() (*RedisService, error) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, redisConnectionErr := redisClient.Ping(ctx).Result()

	if redisConnectionErr != nil {
		return nil, redisConnectionErr
	}
	return &RedisService{
		redisClient: redisClient,
	}, nil
}

func (rc *RedisService) SaveDataToRedis(employeesData []models.Employee) (bool, error) {
	if len(employeesData) == 0 {
		return false, errors.New("Employees list is empty")
	}
	marshalEmployeesData, err := json.Marshal(employeesData)
	if err != nil {
		return false, err
	}
	redisInsertionErr := rc.redisClient.Set(ctx, "employeeData", marshalEmployeesData, 5*time.Minute).Err()

	if redisInsertionErr != nil {
		return false, redisInsertionErr
	}

	return true, nil

}

func (rc *RedisService) GetDataFromRedis() (string, error) {
	employeesData, getErr := rc.redisClient.Get(ctx, "employeeData").Result()

	if getErr != nil {
		return "", getErr
	}
	return employeesData, nil
}
