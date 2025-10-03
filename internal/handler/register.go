package handler

import (
	"go_api/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary      Register a new user
// @Description  Register a new user with the provided information
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RegisterRequest  true  "User registration info"
// @Success      201  {object}  dto.RegisterResponse
// @Failure      400  {object}  map[string]string
// @Failure      409  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var ctx = c.Request.Context()
	var req dto.RegisterRequest

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
	userID, statusCode, err := h.userService.Register(ctx, &req)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, dto.RegisterResponse{
		ID: userID,
	})
}
