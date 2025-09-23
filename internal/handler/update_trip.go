package handler

import (
	"errors"
	"fmt"
	"go_api/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userID not found"})
		return
	}
	fmt.Println(userId)
	id, ok := userId.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid userID type"})
		return
	}
	tripIDstr := c.Param("post_id")
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
}
