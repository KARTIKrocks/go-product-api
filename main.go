package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	controller "goProductAPI/controller"
	"goProductAPI/models"
)

var products = []models.Product{
	{Id: 1, Name: "Product 1", Stock: 10, Price: 2.5},
	{Id: 2, Name: "Product 2", Stock: 5, Price: 3.5},
	{Id: 3, Name: "Product 3", Stock: 7, Price: 5.5},
}

func main() {
	r := gin.Default()
	controller := controller.Init(&products)

	r.GET("/product", controller.ReadProduct)
	r.GET("/product/:id", controller.ReadProductById)
	r.POST("/product", controller.CreateProduct)
	r.PUT("/product/:id", controller.UpdateProductById)
	r.DELETE("/product/:id", controller.DeleteProductById)

	r.Run("localhost:8080")
	fmt.Println("server is running")
}
