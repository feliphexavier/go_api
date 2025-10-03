package handler

import (
	"go_api/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateTrip godoc
// @Summary      Create a new trip
// @Description  Create a new trip with the provided information
// @Tags         Trips
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateOrUpdateTripRequest  true  "Trip creation info"
// @Success      201  {object}  dto.CreateOrUpdateTripResponse
// @Security     BearerAuth
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /trip [post]
func (h *TripHandler) CreateTrip(c *gin.Context) {
	var ctx = c.Request.Context()
	var req dto.CreateOrUpdateTripRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userID not found"})
		return
	}
	id, ok := userId.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid userID type"})
		return
	}
	tripID, statusCode, err := h.tripService.CreateTrip(ctx, &req, id)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, dto.CreateOrUpdateTripResponse{
		ID: tripID,
	})
}
