package main

import (
	"ecommerce-go/controller"
	"ecommerce-go/db"
	"ecommerce-go/repository"
	"ecommerce-go/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// REPOSITORY
	ProductRepository := repository.NewProductRepository(dbConnection)

	// USE CASE (service)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// CONTROLLERS
	ProductController := controller.NewProductController(ProductUseCase)

	// ROUTES
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.GET("/product/:productId", ProductController.GetProductById)

	server.POST("/product", ProductController.CreateProduct)

	server.Run(":8000")
}
