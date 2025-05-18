package productrepo

import (
	"database/sql"
	"encoding/json"
)

type ProductRepo interface {
	CreateProduct(product Product) (Product, error)
	GetProduct(id uint) (Product, error)
	UpdateProduct(id uint, product Product) (Product, error)
	DeleteProduct(id uint) error

	GetProduts() ([]Product, error)
}

type productRepoImpl struct {
	db *sql.DB
}

func New(database *sql.DB) ProductRepo {
	return &productRepoImpl{db: database}
}

func (r *productRepoImpl) CreateProduct(product Product) (Product, error) {
	const query = `
		INSERT INTO products (name, description, price, attributes)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, description, price, attributes, created_at, updated_at`

	var p Product
	var rawAttrs []byte

	err := r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Attributes,
	).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&rawAttrs,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return Product{}, err
	}

	if err := json.Unmarshal(rawAttrs, &p.Attributes); err != nil {
		return Product{}, err
	}

	return p, nil
}

// GetProduct возвращает продукт по id
func (r *productRepoImpl) GetProduct(id uint) (Product, error) {
	const query = `
		SELECT id, name, description, price, attributes, created_at, updated_at
		FROM products
		WHERE id = $1`

	var p Product
	var rawAttrs []byte

	err := r.db.QueryRow(query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&rawAttrs,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return Product{}, err
	}

	if err := json.Unmarshal(rawAttrs, &p.Attributes); err != nil {
		return Product{}, err
	}

	return p, nil
}

// UpdateProduct обновляет все поля продукта и возвращает обновлённый объект
func (r *productRepoImpl) UpdateProduct(id uint, product Product) (Product, error) {
	const query = `
		UPDATE products
		SET
			name        = $1,
			description = $2,
			price       = $3,
			attributes  = $4,
			updated_at  = NOW()
		WHERE id = $5
		RETURNING id, name, description, price, attributes, created_at, updated_at`

	var p Product
	var rawAttrs []byte

	err := r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Attributes,
		id,
	).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&rawAttrs,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return Product{}, err
	}

	if err := json.Unmarshal(rawAttrs, &p.Attributes); err != nil {
		return Product{}, err
	}

	return p, nil
}

func (r *productRepoImpl) DeleteProduct(id uint) error {
	const query = `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *productRepoImpl) GetProduts() ([]Product, error) {
	const query = `
		SELECT id, name, description, price, attributes, created_at, updated_at
		FROM products`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Product
	for rows.Next() {
		var p Product
		var rawAttrs []byte

		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&rawAttrs,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(rawAttrs, &p.Attributes); err != nil {
			return nil, err
		}

		list = append(list, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}
