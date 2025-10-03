package handler

import (
	_ "go_api/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetTripByID godoc
// @Summary      Get trip by ID
// @Description  Retrieve trip details by trip ID
// @Tags         Trips
// @Accept       json
// @Produce      json
// @Param        trip_id   path      string  true  "trip_id"
// @Security     BearerAuth
// @Success      200  {object}  dto.GetTripResponse
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /trip/{trip_id} [get]
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
