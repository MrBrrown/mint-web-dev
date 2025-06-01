package orderrepo

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type OrderRepo interface {
	CreateOrder(order Order) (Order, error)
	GetOrder(id uint) (Order, error)
	UpdateOrder(id uint, order Order) (Order, error)
	UpdateOrderStatus(id uint, status string) (Order, error)
	DeleteOrder(id uint) error
	ListOrders() ([]Order, error)
}

type orderRepoImpl struct {
	db *sql.DB
}

func New(db *sql.DB) OrderRepo {
	return &orderRepoImpl{db: db}
}

func (r *orderRepoImpl) CreateOrder(order Order) (Order, error) {
	const query = `
		INSERT INTO orders (total_price, items, user_info, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, total_price, items, user_info, status, created_at, updated_at`

	var o Order
	var itemsJSON, userInfoJSON []byte

	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return Order{}, fmt.Errorf("failed to marshal items: %w", err)
	}

	userInfoJSON, err = json.Marshal(order.UserInfo)
	if err != nil {
		return Order{}, fmt.Errorf("failed to marshal user info: %w", err)
	}

	err = r.db.QueryRow(query,
		order.TotalPrice,
		itemsJSON,
		userInfoJSON,
		order.Status,
	).Scan(
		&o.ID,
		&o.TotalPrice,
		&itemsJSON,
		&userInfoJSON,
		&o.Status,
		&o.CreatedAt,
		&o.UpdatedAt,
	)
	if err != nil {
		return Order{}, err
	}

	_ = json.Unmarshal(itemsJSON, &o.Items)
	_ = json.Unmarshal(userInfoJSON, &o.UserInfo)

	return o, nil
}

func (r *orderRepoImpl) GetOrder(id uint) (Order, error) {
	const query = `
		SELECT id, total_price, items, user_info, status, created_at, updated_at
		FROM orders
		WHERE id = $1`

	var o Order
	var rawItems, rawUserInfo []byte

	err := r.db.QueryRow(query, id).Scan(
		&o.ID,
		&o.TotalPrice,
		&rawItems,
		&rawUserInfo,
		&o.Status,
		&o.CreatedAt,
		&o.UpdatedAt,
	)
	if err != nil {
		return Order{}, err
	}

	if err := json.Unmarshal(rawItems, &o.Items); err != nil {
		return Order{}, err
	}
	if err := json.Unmarshal(rawUserInfo, &o.UserInfo); err != nil {
		return Order{}, err
	}

	return o, nil
}

func (r *orderRepoImpl) UpdateOrder(id uint, order Order) (Order, error) {
	const query = `
		UPDATE orders
		SET total_price = $1, items = $2, user_info = $3, status = $4, updated_at = NOW()
		WHERE id = $5
		RETURNING id, total_price, items, user_info, status, created_at, updated_at`

	var o Order
	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return Order{}, fmt.Errorf("failed to marshal items: %w", err)
	}

	userInfoJSON, err := json.Marshal(order.UserInfo)
	if err != nil {
		return Order{}, fmt.Errorf("failed to marshal user info: %w", err)
	}

	var rawItems, rawUserInfo []byte

	err = r.db.QueryRow(query,
		order.TotalPrice,
		itemsJSON,
		userInfoJSON,
		order.Status,
		id,
	).Scan(
		&o.ID,
		&o.TotalPrice,
		&rawItems,
		&rawUserInfo,
		&o.Status,
		&o.CreatedAt,
		&o.UpdatedAt,
	)
	if err != nil {
		return Order{}, err
	}

	if err := json.Unmarshal(rawItems, &o.Items); err != nil {
		return Order{}, err
	}
	if err := json.Unmarshal(rawUserInfo, &o.UserInfo); err != nil {
		return Order{}, err
	}

	return o, nil
}

func (r *orderRepoImpl) UpdateOrderStatus(id uint, status string) (Order, error) {
	const query = `
		UPDATE orders
		SET status = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING id, total_price, items, user_info, status, created_at, updated_at`

	var o Order
	var itemsJSON, userInfoJSON []byte

	err := r.db.QueryRow(query, status, id).Scan(
		&o.ID,
		&o.TotalPrice,
		&itemsJSON,
		&userInfoJSON,
		&o.Status,
		&o.CreatedAt,
		&o.UpdatedAt,
	)
	if err != nil {
		return Order{}, err
	}

	if err := json.Unmarshal(itemsJSON, &o.Items); err != nil {
		return Order{}, err
	}
	if err := json.Unmarshal(userInfoJSON, &o.UserInfo); err != nil {
		return Order{}, err
	}

	return o, nil
}

func (r *orderRepoImpl) DeleteOrder(id uint) error {
	const query = `DELETE FROM orders WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *orderRepoImpl) ListOrders() ([]Order, error) {
	const query = `
		SELECT id, total_price, items, user_info, status, created_at, updated_at
		FROM orders`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]Order, 0)
	for rows.Next() {
		var o Order
		var rawItems, rawUserInfo []byte

		if err := rows.Scan(
			&o.ID,
			&o.TotalPrice,
			&rawItems,
			&rawUserInfo,
			&o.Status,
			&o.CreatedAt,
			&o.UpdatedAt,
		); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(rawItems, &o.Items); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(rawUserInfo, &o.UserInfo); err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	return orders, rows.Err()
}
