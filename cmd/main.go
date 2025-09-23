package main

import (
	"context"
	"fmt"
	"go_api/internal/config"
	"go_api/internal/dto"
	tripHandler "go_api/internal/handler"
	userHandler "go_api/internal/handler"
	"go_api/internal/model"
	tripRepo "go_api/internal/repository"
	userRepo "go_api/internal/repository"
	tripService "go_api/internal/service"
	userService "go_api/internal/service"
	"go_api/pkg"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()
	validate := validator.New()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := pkg.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	userRepo := userRepo.NewRepository(db)
	tripRepo := tripRepo.NewTripRepository(db)

	userService := userService.NewService(cfg, userRepo)
	tripService := tripService.NewTripService(cfg, tripRepo)

	userHandler := userHandler.NewHandler(r, validate, userService)
	tripHandler := tripHandler.NewTripHandler(r, validate, tripService)

	userHandler.RouteList(cfg.SecretJWT)
	tripHandler.RouteList(cfg.SecretJWT)
	id := "9dabe229-9751-4a95-9af7-c888b0199c68"
	uid := "75b54fb6-a55b-4c3c-8be2-65382436fc09"
	rp := tripRepo.UpdateTrip(context.Background(), &model.TripModel{Title: "a", Description: "batora", Start_date: "2020-05-05", End_date: "2020-05-05"}, uuid.MustParse(id))
	sv, n, err := tripService.UpdateTrip(context.Background(), &dto.CreateOrUpdateTripRequest{Title: "a", Description: "batora", Start_Date: "2020-04-04", End_Date: "2020-04-04"}, uuid.MustParse(id), uuid.MustParse(uid))
	fmt.Println(sv, n)
	fmt.Println(rp)
	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
