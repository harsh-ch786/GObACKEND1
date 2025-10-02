package repository

import (
	"context"

	"github.com/harsh-karwar/udhaar-tracker-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUdhaar(ctx context.Context, userID string, udhaarID string, updatedUdhaar models.Udhaar) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": udhaarID, "userId": userID}
	update := bson.M{
		"$set": bson.M{
			"friendName":  updatedUdhaar.FriendName,
			"amount":      updatedUdhaar.Amount,
			"description": updatedUdhaar.Description,
			"status":      updatedUdhaar.Status,
			"dueDate":     updatedUdhaar.DueDate,
		},
	}

	// Find the document and update it.
	result, err := udhaarCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}