package routes

import (
	"github.com/Gambi18/Quizzo/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/user/:id", controller.GetUser)
	router.GET("/user/", controller.GetUsers)
	router.POST("/user/", controller.CreateUser)
	router.DELETE("/user/:id", controller.DeleteUser)
	router.PUT("/user/:id", controller.UpdateUser)
}

func ScoreRoute(router *gin.Engine) {
	router.GET("/score/", controller.GetScores)
    router.GET("/score/:id", controller.GetScore)
    // router.POST("/score/", controller.CreateScore)
    router.PUT("/score/:id", controller.UpdateScore)
}
