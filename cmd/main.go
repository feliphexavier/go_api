package main

import (
	"fmt"
	_ "go_api/cmd/docs"
	"go_api/internal/config"
	pictureHandler "go_api/internal/handler"
	tripHandler "go_api/internal/handler"
	userHandler "go_api/internal/handler"
	pictureRepo "go_api/internal/repository"
	tripRepo "go_api/internal/repository"
	userRepo "go_api/internal/repository"
	pictureService "go_api/internal/service"
	tripService "go_api/internal/service"
	userService "go_api/internal/service"
	"go_api/pkg"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Photo Album API
// @version         1.0
// @description     This is a sample server for a trip booking service.
// @termsOfService  http://swagger.io/terms/

//@securityDefinitions.apikey BearerAuth
//@in header
//@name Authorization
//@description Type "Bearer" followed by a space and JWT token.

// @host 	localhost:8080
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
	pictureRepo := pictureRepo.NewPictureRepository(db)

	userService := userService.NewService(cfg, userRepo)
	tripService := tripService.NewTripService(cfg, tripRepo, pictureRepo)
	pictureService := pictureService.NewPictureService(cfg, pictureRepo, tripRepo)

	userHandler := userHandler.NewHandler(r, validate, userService)
	tripHandler := tripHandler.NewTripHandler(r, validate, tripService)
	pictureHandler := pictureHandler.NewPictureHandler(r, pictureService)

	pictureHandler.RouteList(cfg.SecretJWT)
	userHandler.RouteList(cfg.SecretJWT)
	tripHandler.RouteList(cfg.SecretJWT)
	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(server)
}
