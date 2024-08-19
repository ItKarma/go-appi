package main

import (
	"go-appi/controller"
	"go-appi/db"
	"go-appi/repository"
	"go-appi/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRespository(dbConnection)

	ProductUsecase := usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":3000")
}
