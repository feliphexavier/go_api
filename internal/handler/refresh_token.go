package handler

import (
	"go_api/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RefreshToken godoc
// @Summary      Refresh JWT Token
// @Description  Refresh the JWT token using a valid refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RefreshTokenRequest  true  "Refresh token info"
// @Success      200  {object}  dto.RefreshTokenResponse
// @Security     BearerAuth
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /auth/refresh [post]
func (h *Handler) RefreshToken(c *gin.Context) {
	var ctx = c.Request.Context()
	var req dto.RefreshTokenRequest

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
	id, ok := userId.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid userID type"})
		return
	}
	token, refreshToken, statusCode, err := h.userService.RefreshToken(ctx, &req, id)

	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, dto.RefreshTokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}
