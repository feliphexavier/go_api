package handler

import (
	"errors"
	"go_api/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *PictureHandler) CreatePicture(c *gin.Context) {
	var ctx = c.Request.Context()
	var req dto.CreatePictureRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors.New("Invalid request body"),
		})
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
	PictureID, err := h.pictureService.CreatePicture(ctx, &req, tripID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, dto.CreatePictureReponse{
		ID: PictureID,
	})
	return
}
