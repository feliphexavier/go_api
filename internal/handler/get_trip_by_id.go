package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *TripHandler) GetTripByID(c *gin.Context) {
	ctx := c.Request.Context()
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userID not found"})
		return
	}
	userID, ok := userId.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid userID type"})
		return
	}
	tripIDParam := c.Param("trip_id")
	tripID, err := uuid.Parse(tripIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trip ID"})
		return
	}
	trip, statusCode, err := h.tripService.GetTripByID(ctx, tripID, userID)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	if trip == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trip not found"})
		return
	}
	c.JSON(http.StatusOK, trip)
	return
}
