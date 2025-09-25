package handler

import (
	"go_api/internal/middleware"
	"go_api/internal/service"

	"github.com/gin-gonic/gin"
)

type PictureHandler struct {
	api            *gin.Engine
	pictureService service.PictureService
}

func NewPictureHandler(api *gin.Engine, pictureService service.PictureService) *PictureHandler {
	return &PictureHandler{
		api:            api,
		pictureService: pictureService,
	}
}
func (h *PictureHandler) RouteList(secretKey string) {
	routeAuth := h.api.Group("/picture")
	routeAuth.Use(middleware.AuthMiddleware(secretKey))
	routeAuth.POST("/:trip_id", h.CreatePicture)
}
