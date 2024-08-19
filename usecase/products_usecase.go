package usecase

import (
	"go-appi/model"
	"go-appi/repository"
)

type ProductUsecase struct {
	//Repository

	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil
}

func (pu *ProductUsecase) GetProducByID(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductByID(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
