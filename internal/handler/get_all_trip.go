package handler

import (
	"go_api/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllTrips godoc
// @Summary      Get all trips
// @Description  Retrieve a paginated list of all trips for the authenticated user
// @Tags         Trips
// @Accept       json
// @Produce      json
// @Param        page   query      int  false  "Page number"  default(1)
// @Param        limit  query      int  false  "Number of items per page"  default(10)
// @Security     BearerAuth
// @Success      200  {object}  dto.GetAllTripsResponse
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /trip [get]
func (h *TripHandler) GetAllTrips(c *gin.Context) {
	ctx := c.Request.Context()
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(400, gin.H{"error": "Invalid page parameter"})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.JSON(400, gin.H{"error": "Invalid limit parameter"})
		return
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
	trips, statusCode, err := h.tripService.GetAllTrip(ctx, &dto.GetAllTripsRequest{
		Page:  page,
		Limit: limit,
	}, id)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, trips)
	return
}
