package main

import (
	"shop/controller"
	"shop/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "PATCH", "GET", "DELETE", "POST"},
	}))
	router.GET("/product", controller.GetProduct)
	router.POST("/product", controller.CreateProduct)
	router.PATCH("/product/:id", controller.UpdateProduct)
	router.DELETE("/product/:id", controller.DeleteProduct)

	router.Run()
}
