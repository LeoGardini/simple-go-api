package routes

import (
	"github.com/LeoGardini/simple-go-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")

	books := main.Group("books")

	books.GET("/", controllers.ListBooks)
	books.GET("/:id", controllers.ShowBook)

	books.POST("/", controllers.CreateBook)

	return router
}
