package controller

import (
	"go-appi/model"
	"go-appi/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	//Usecase
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUseCase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products, err := p.productUseCase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *ProductController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")

	if id == "" {
		response := model.Response{
			Message: "Id do produto é obrigatorio",
		}
		ctx.JSON(http.StatusBadRequest, response)
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id deve ser um numero !",
		}
		ctx.JSON(http.StatusBadRequest, response)
	}

	product, err := p.productUseCase.GetProducByID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não foi encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
	}

	ctx.JSON(http.StatusCreated, product)
}
