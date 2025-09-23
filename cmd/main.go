package main

import (
	"fmt"
	"go_api/internal/config"
	tripHandler "go_api/internal/handler"
	userHandler "go_api/internal/handler"
	tripRepo "go_api/internal/repository"
	userRepo "go_api/internal/repository"
	tripService "go_api/internal/service"
	userService "go_api/internal/service"
	"go_api/pkg"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	/*id := "1909bd03-68f1-48f5-90fa-00a6ffb29c71"
	uid := "75b54fb6-a55b-4c3c-8be2-65382436fc09"
	fmt.Println(tripService.DeleteTrip(context.Background(), uuid.MustParse(id), uuid.MustParse(uid)))*/
	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
