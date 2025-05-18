package productusecase

import (
	"errors"
	productrepo "marketapi/products/internal/repositories/product_repo"
)

type ProductUseCase struct {
	repo productrepo.ProductRepo
}

func New(r productrepo.ProductRepo) *ProductUseCase {
	return &ProductUseCase{repo: r}
}

func (uc *ProductUseCase) CreateProduct(p productrepo.Product) (productrepo.Product, error) {
	if p.Name == "" {
		return productrepo.Product{}, errors.New("name is required")
	}
	if p.Price < 0 {
		return productrepo.Product{}, errors.New("price must be non-negative")
	}

	return uc.repo.CreateProduct(p)
}

func (uc *ProductUseCase) GetProduct(id uint) (productrepo.Product, error) {
	return uc.repo.GetProduct(id)
}

func (uc *ProductUseCase) UpdateProduct(id uint, p productrepo.Product) (productrepo.Product, error) {
	if p.Name == "" {
		return productrepo.Product{}, errors.New("name is required")
	}
	if p.Price < 0 {
		return productrepo.Product{}, errors.New("price must be non-negative")
	}

	return uc.repo.UpdateProduct(id, p)
}

func (uc *ProductUseCase) DeleteProduct(id uint) error {
	return uc.repo.DeleteProduct(id)
}

func (uc *ProductUseCase) ListProducts() ([]productrepo.Product, error) {
	return uc.repo.GetProduts()
}
