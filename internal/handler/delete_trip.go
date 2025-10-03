package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DeleteTrip godoc
// @Summary      Delete trip
// @Description  Delete trip by trip ID
// @Tags         Trips
// @Accept       json
// @Produce      json
// @Param        trip_id   path      string  true  "trip_id"
// @Security     BearerAuth
// @Success      204  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /trip/{trip_id} [delete]
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

	tripIDstr := c.Param("trip_id")
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
	return
}
