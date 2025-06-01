package orderusecase

import (
	"errors"
	"time"

	orderrepo "marketapi/orders/internal/repositories/order_repo"
)

type CreateOrderRequest struct {
	Items    []OrderItem    `json:"items"`
	UserInfo map[string]any `json:"user_info"`
	Status   string         `json:"status"`
}

type OrderItem struct {
	ID       uint `json:"id"`
	Quantity uint `json:"quantity"`
}

type OrderUseCase struct {
	repo orderrepo.OrderRepo
}

func New(r orderrepo.OrderRepo) *OrderUseCase {
	return &OrderUseCase{repo: r}
}

func (uc *OrderUseCase) CreateOrder(req CreateOrderRequest) (orderrepo.Order, error) {
	o := orderrepo.Order{}
	return uc.repo.CreateOrder(o)
}

func (uc *OrderUseCase) GetOrder(id uint) (orderrepo.Order, error) {
	return uc.repo.GetOrder(id)
}

func (uc *OrderUseCase) UpdateOrder(id uint, o orderrepo.Order) (orderrepo.Order, error) {
	if o.TotalPrice <= 0 {
		return orderrepo.Order{}, errors.New("order total price must be positive")
	}
	if len(o.Items) == 0 {
		return orderrepo.Order{}, errors.New("order must contain at least one item")
	}
	o.UpdatedAt = time.Now()
	return uc.repo.UpdateOrder(id, o)
}

func (uc *OrderUseCase) UpdateStatus(id uint, status string) (orderrepo.Order, error) {
	if status == "" {
		return orderrepo.Order{}, errors.New("status cannot be empty")
	}
	return uc.repo.UpdateOrderStatus(id, status)
}

func (uc *OrderUseCase) DeleteOrder(id uint) error {
	return uc.repo.DeleteOrder(id)
}

func (uc *OrderUseCase) ListOrders() ([]orderrepo.Order, error) {
	return uc.repo.ListOrders()
}
