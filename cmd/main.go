package main

import (
	"fmt"
	"go_api/internal/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
