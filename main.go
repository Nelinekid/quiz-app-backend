package main

import (
	"github.com/Gambi18/Quizzo/config"
	"github.com/Gambi18/Quizzo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	routes.ScoreRoute(router)

	router.Run(":8081")
}
