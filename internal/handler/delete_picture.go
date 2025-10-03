package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DeletePicture godoc
// @Summary      Delete picture
// @Description  Delete picture by picture ID
// @Tags         Pictures
// @Accept       json
// @Produce      json
// @Param        picture_id   path      string  true  "picture_id"
// @Security     BearerAuth
// @Success      204  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /picture/{picture_id} [delete]
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
