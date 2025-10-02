package repository

import (
	"context"

	"github.com/harsh-karwar/udhaar-tracker-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)
func GetUdhaarsByUserID(ctx context.Context, userID string) ([]models.Udhaar, error) {
	var udhaars []models.Udhaar
	filter := bson.M{"userId": userID}
	cursor, err := udhaarCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &udhaars); err != nil {
		return nil, err
	}

	return udhaars, nil
}
func GetUdhaarByID(ctx context.Context, userID string, udhaarID string) (models.Udhaar, error) {
	var udhaar models.Udhaar
	filter := bson.M{"_id": udhaarID, "userId": userID}
	err := udhaarCollection.FindOne(ctx, filter).Decode(&udhaar)
	if err != nil {
		return udhaar, err
	}

	return udhaar, nil
}