package handler

import (
	"go_api/internal/middleware"
	sevice "go_api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api         *gin.Engine
	validate    *validator.Validate
	userService sevice.UserService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, userService sevice.UserService) *Handler {
	return &Handler{
		api:         api,
		validate:    validate,
		userService: userService,
	}
}
func (h *Handler) RouteList(secretkey string) {
	authRoute := h.api.Group("/auth")
	authRoute.POST("/register", h.Register)
	authRoute.POST("/login", h.Login)
	refreshRoute := h.api.Group("/auth")
	refreshRoute.Use(middleware.AuthRefreshTokenMiddleware(secretkey))
	refreshRoute.POST("/refresh", h.RefreshToken)
}
