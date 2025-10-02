package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh-karwar/udhaar-tracker-backend/internal/models"
	"github.com/harsh-karwar/udhaar-tracker-backend/internal/repository"
)
func UpdateUdhaar(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is required"})
		return
	}
	udhaarID := c.Param("id")
	var updatedUdhaar models.Udhaar
	if err := c.ShouldBindJSON(&updatedUdhaar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}
	result, err := repository.UpdateUdhaar(c.Request.Context(), userID, udhaarID, updatedUdhaar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update udhaar"})
		return
	}
	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Udhaar not found or you do not have permission to edit it"})
		return
	}
	go repository.ClearUserUdhaarsCache(c.Request.Context(), userID)
	c.JSON(http.StatusOK, gin.H{"message": "Udhaar updated successfully"})
}
