package productrepo

import "database/sql"

type ProductRepo interface {
	CreateProduct(db *sql.DB, product Product) (Product, error)
	GetProduct(db *sql.DB, id uint) (Product, error)
	UpdateProduct(db *sql.DB, id uint) (Product, error)
	DeleteProduct(db *sql.DB, id uint) error

	GetProduts(db *sql.DB) ([]Product, error)
}

type productRepoImpl struct{}

func New() ProductRepo {
	return &productRepoImpl{}
}

func (r *productRepoImpl) CreateProduct(db *sql.DB, product Product) (Product, error) {
	panic("not implemented")
}

func (r *productRepoImpl) GetProduct(db *sql.DB, id uint) (Product, error) {
	panic("not implemented")
}

func (r *productRepoImpl) UpdateProduct(db *sql.DB, id uint) (Product, error) {
	panic("not implemented")
}

func (r *productRepoImpl) DeleteProduct(db *sql.DB, id uint) error {
	panic("not implemented")
}

func (r *productRepoImpl) GetProduts(db *sql.DB) ([]Product, error) {
	panic("not implemented")
}
