package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/harsh-karwar/udhaar-tracker-backend/internal/models"
	"github.com/harsh-karwar/udhaar-tracker-backend/internal/repository"
)
func CreateUdhaar(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is required"})
		return
	}

	var newUdhaar models.Udhaar
	if err := c.ShouldBindJSON(&newUdhaar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUdhaar.ID = uuid.New().String()
	newUdhaar.UserID = userID
	newUdhaar.Status = "pending"
	newUdhaar.CreatedAt = time.Now()
	newUdhaar.DueDate = time.Now().Add(30 * 24 * time.Hour)

	_, err := repository.CreateUdhaar(c.Request.Context(), &newUdhaar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create udhaar"})
		return
	}
	go repository.ClearUserUdhaarsCache(c.Request.Context(), userID)

	c.JSON(http.StatusCreated, newUdhaar)
}