package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *PictureHandler) DeletePicture(c *gin.Context) {
	var ctx = c.Request.Context()

	pictureIDstr := c.Param("picture_id")
	pictureID, err := uuid.Parse(pictureIDstr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors.New("Parsing error"),
		})
		return
	}
	err = h.pictureService.DeletePicture(ctx, pictureID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "deleted"})
	return
}
