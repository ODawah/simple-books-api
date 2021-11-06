package main

import (
	"example.com/m/v2/controllers"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("books/:id", controllers.UpdateBook)
	r.DELETE("books/:id", controllers.DeleteBook)

	models.ConnectDatabase()

	r.Run()
}
