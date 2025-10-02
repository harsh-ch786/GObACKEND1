package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/harsh-karwar/udhaar-tracker-backend/internal/database"
	"github.com/harsh-karwar/udhaar-tracker-backend/internal/models"
)

const (
	udhaarListCacheDuration = 10 * time.Minute
)

func getUserUdhaarsCacheKey(userID string) string {
	return fmt.Sprintf("udhaars:%s", userID)
}


func SetUserUdhaarsCache(ctx context.Context, userID string, udhaars []models.Udhaar) error {
	key := getUserUdhaarsCacheKey(userID)
	data, err := json.Marshal(udhaars)
	if err != nil {
		return err
	}
	err = database.Redis.Set(ctx, key, data, udhaarListCacheDuration).Err()
	return err
}


func GetUserUdhaarsCache(ctx context.Context, userID string) ([]models.Udhaar, error) {
	key := getUserUdhaarsCacheKey(userID)


	result, err := database.Redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var udhaars []models.Udhaar
	err = json.Unmarshal([]byte(result), &udhaars)
	if err != nil {
		return nil, err
	}

	return udhaars, nil
}
func ClearUserUdhaarsCache(ctx context.Context, userID string) error {
	key := getUserUdhaarsCacheKey(userID)
	err := database.Redis.Del(ctx, key).Err()
	return err
}