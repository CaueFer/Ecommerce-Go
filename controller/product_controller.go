package controller

import (
	"ecommerce-go/model"
	"ecommerce-go/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUseCase usecase.ProductUseCase
}

func NewProductController(useCase usecase.ProductUseCase) *ProductController {
	return &ProductController{
		ProductUseCase: useCase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products, err := p.ProductUseCase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)

	if(err != nil){
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.ProductUseCase.CreateProduct(product)

	if(err != nil){
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}