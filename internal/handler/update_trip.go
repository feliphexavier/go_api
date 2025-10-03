package handler

import (
	"errors"
	"fmt"
	"go_api/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UpdateTrip godoc
// @Summary      Update trip
// @Description  Update trip details by trip ID
// @Tags         Trips
// @Accept       json
// @Produce      json
// @Param        trip_id   path      string  true  "trip_id"
// @Param        trip      body      dto.CreateOrUpdateTripRequest  true  "Trip details"
// @Security     BearerAuth
// @Success      200  {object}  dto.CreateOrUpdateTripResponse
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /trip/{trip_id} [put]
func (h *TripHandler) UpdateTrip(c *gin.Context) {
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
	fmt.Println(userId)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userID not found"})
		return
	}
	id, ok := userId.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid userID type"})
		return
	}
	tripIDstr := c.Param("trip_id")
	tripID, err := uuid.Parse(tripIDstr)
	if err != nil {
		c.JSON(http.StatusContinue, gin.H{
			"message": errors.New("Parsing error"),
		})
		return
	}
	tripID, statusCode, err := h.tripService.UpdateTrip(ctx, &req, tripID, id)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, dto.CreateOrUpdateTripResponse{
		ID: tripID,
	})
	return
}
