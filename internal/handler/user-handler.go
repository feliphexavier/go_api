package handler

import (
	sevice "go_api/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	api         *gin.Engine
	userService sevice.UserService
}

func NewHandler(api *gin.Engine, userService sevice.UserService) *Handler {
	return &Handler{
		api:         api,
		userService: userService,
	}
}
func (h *Handler) RouteList() {
	authRoute := h.api.Group("/auth")
	authRoute.POST("/register", h.Register)
}
