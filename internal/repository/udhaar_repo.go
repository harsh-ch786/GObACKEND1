package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/harsh-karwar/udhaar-tracker-backend/internal/database"
	"github.com/harsh-karwar/udhaar-tracker-backend/internal/models"
)

var udhaarCollection *mongo.Collection
func InitRepository() {
	udhaarCollection = database.GetCollection("udhaars")
}
func CreateUdhaar(ctx context.Context, udhaar *models.Udhaar) (*mongo.InsertOneResult, error) {
	result, err := udhaarCollection.InsertOne(ctx, udhaar)
	if err != nil {
		return nil, err
	}
	return result, nil
}
