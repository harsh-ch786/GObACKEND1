package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh-karwar/udhaar-tracker-backend/internal/repository"
)

// DeleteUdhaar handles the DELETE /udhaars/:id endpoint.
func DeleteUdhaar(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is required"})
		return
	}
	udhaarID := c.Param("id")

	result, err := repository.DeleteUdhaar(c.Request.Context(), userID, udhaarID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete udhaar"})
		return
	}
	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Udhaar not found or you do not have permission to delete it"})
		return
	}

	go repository.ClearUserUdhaarsCache(c.Request.Context(), userID)

	c.JSON(http.StatusOK, gin.H{"message": "Udhaar deleted successfully"})
}

