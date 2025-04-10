package main

import (
	"inventory/handlers"

	"github.com/gin-gonic/gin"
)

var (
	product_handler *handlers.ProductHandler
)

func main() {
	router := gin.Default()

	product_handler = handlers.NewProductHandler()

	router.POST("/products", product_handler.SaveProduct)

	router.GET("/products/:id", product_handler.GetProduct)

	router.PATCH("/products/:id", product_handler.UpdateProduct)

	router.DELETE("/products/:id", product_handler.DeleteProduct)

	router.GET("/products", product_handler.ListProducts)

	router.Run("0.0.0.0:8082")
}
