package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/harsh-karwar/udhaar-tracker-backend/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetUdhaars(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is required"})
		return
	}


	cachedUdhaars, err := repository.GetUserUdhaarsCache(c.Request.Context(), userID)
	if err == nil {

		log.Println("Cache HIT for user:", userID)
		c.JSON(http.StatusOK, cachedUdhaars)
		return
	}


	if err != redis.Nil {
		log.Printf("Redis error for user %s: %v", userID, err)
	}

	log.Println("Cache MISS for user:", userID)
	udhaars, err := repository.GetUdhaarsByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch udhaars"})
		return
	}

	if udhaars == nil {
		c.JSON(http.StatusOK, []interface{}{})
		return
	}

	go repository.SetUserUdhaarsCache(c.Request.Context(), userID, udhaars)
	c.JSON(http.StatusOK, udhaars)
}

func GetUdhaarByID(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is required"})
		return
	}

	udhaarID := c.Param("id")

	udhaar, err := repository.GetUdhaarByID(c.Request.Context(), userID, udhaarID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Udhaar not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch udhaar"})
		return
	}

	c.JSON(http.StatusOK, udhaar)
}

