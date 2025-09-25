package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *TripHandler) DeleteTrip(c *gin.Context) {
	var ctx = c.Request.Context()

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

	tripIDstr := c.Param("post_id")
	tripID, err := uuid.Parse(tripIDstr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors.New("Parsing error"),
		})
		return
	}
	_, err = h.tripService.DeleteTrip(ctx, tripID, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "deleted"})
}
