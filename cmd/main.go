package main

import (
	"github.com/gin-gonic/gin"
	"ecommerce-go/controller"
	"ecommerce-go/usecase"
	"ecommerce-go/db"
	"ecommerce-go/repository"
)

func main() {
	server := gin.Default()
	
	dbConnection, err := db.ConnectDB()
	if(err != nil){
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

	server.POST("/product", ProductController.CreateProduct)

	server.Run(":8000")
}
