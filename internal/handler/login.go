package handler

import (
	"go_api/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var ctx = c.Request.Context()
	var req dto.LoginRequest

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
	token, refresh_token, statusCode, err := h.userService.Login(ctx, &req)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, dto.LoginResponse{
		Token:        token,
		RefreshToken: refresh_token,
	})
}
