package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/harsh-karwar/udhaar-tracker-backend/internal/database"
	"github.com/harsh-karwar/udhaar-tracker-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection
func InitUserRepository() {
	userCollection = database.GetCollection("users")
}
func CreateUser(ctx context.Context, payload models.SignupPayload) (*mongo.InsertOneResult, error) {
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": payload.Email})
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("user with this email already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := models.User{
		ID:       uuid.New().String(),
		Email:    payload.Email,
		Password: string(hashedPassword),
	}
	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	filter := bson.M{"email": email}

	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user,nil
}