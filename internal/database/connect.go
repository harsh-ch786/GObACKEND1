package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Mongo *mongo.Client
	Redis *redis.Client
	Ctx   = context.Background()
)

func ConnectDB() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in .env file")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()


client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}

if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Could not ping MongoDB: %v", err)
}

	
Mongo = client
fmt.Println("Connected to MongoDB!")
redisAddr := os.Getenv("REDIS_ADDR")
if redisAddr == "" {
	  redisAddr = "redis:6379"
}

rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0, 
})


	_, err = rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}


Redis = rdb
fmt.Println("Connected to Redis!")
}


func GetCollection(collectionName string) *mongo.Collection {
	return Mongo.Database("udhaar_db").Collection(collectionName)
}