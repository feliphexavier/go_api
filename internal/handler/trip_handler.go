package handler

import (
	"go_api/internal/middleware"
	"go_api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TripHandler struct {
	api         *gin.Engine
	validate    *validator.Validate
	tripService service.TripService
}

func NewTripHandler(api *gin.Engine, validate *validator.Validate, tripService service.TripService) *TripHandler {
	return &TripHandler{
		api:         api,
		validate:    validate,
		tripService: tripService}
}
func (h *TripHandler) RouteList(secretKey string) {
	routeAuth := h.api.Group("/trip")
	routeAuth.Use(middleware.AuthMiddleware(secretKey))
	routeAuth.GET("/:trip_id", h.GetTripByID)
	routeAuth.GET("", h.GetAllTrips)
	routeAuth.POST("", h.CreateTrip)
	routeAuth.PUT("/:trip_id/update", h.UpdateTrip)
	routeAuth.DELETE("/:trip_id/delete", h.DeleteTrip)
}
