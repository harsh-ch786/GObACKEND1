package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteUdhaar(ctx context.Context, userID string, udhaarID string) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": udhaarID, "userId": userID}
	result, err := udhaarCollection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
