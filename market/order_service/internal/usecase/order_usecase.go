package orderusecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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

type Product struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"desc"`
	Price       float64        `json:"price"`
	Attributes  map[string]any `json:"attribs"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"update_at"`
}

type OrderUseCase struct {
	repo orderrepo.OrderRepo
}

func New(r orderrepo.OrderRepo) *OrderUseCase {
	return &OrderUseCase{repo: r}
}

func (uc *OrderUseCase) CreateOrder(req CreateOrderRequest) (orderrepo.Order, error) {
	log.Printf("[CreateOrder] Received request: %+v", req)

	var productIDs []uint
	quantities := make(map[uint]uint)
	for _, item := range req.Items {
		productIDs = append(productIDs, item.ID)
		quantities[item.ID] = item.Quantity
	}

	log.Printf("[CreateOrder] Extracted product IDs: %v", productIDs)

	payload, err := json.Marshal(productIDs)
	if err != nil {
		log.Printf("[CreateOrder] Failed to marshal product IDs: %v", err)
		return orderrepo.Order{}, fmt.Errorf("failed to marshal product IDs: %w", err)
	}

	log.Printf("[CreateOrder] Sending request to product service")

	resp, err := http.Post("http://gateway:8080/products/batch", "application/json", bytes.NewReader(payload))
	if err != nil {
		log.Printf("[CreateOrder] Failed to request products: %v", err)
		return orderrepo.Order{}, fmt.Errorf("failed to request products: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("[CreateOrder] Product service returned error: %s", body)
		return orderrepo.Order{}, fmt.Errorf("product service error: %s", body)
	}

	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		log.Printf("[CreateOrder] Failed to decode products: %v", err)
		return orderrepo.Order{}, fmt.Errorf("failed to decode products: %w", err)
	}

	log.Printf("[CreateOrder] Received products: %+v", products)

	var (
		items      []map[string]any
		totalPrice float64
	)

	for _, p := range products {
		qty := quantities[p.ID]
		subtotal := p.Price * float64(qty)

		item := map[string]any{
			"name":        p.Name,
			"description": p.Description,
			"price":       p.Price,
			"quantity":    qty,
		}

		items = append(items, item)
		totalPrice += subtotal
	}

	log.Printf("[CreateOrder] Total price calculated: %.2f", totalPrice)

	order := orderrepo.Order{
		TotalPrice: totalPrice,
		Items:      items,
		UserInfo:   req.UserInfo,
		Status:     req.Status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	log.Printf("[CreateOrder] Final order to save: %+v", order)

	createdOrder, err := uc.repo.CreateOrder(order)
	if err != nil {
		log.Printf("[CreateOrder] Failed to save order: %v", err)
		return orderrepo.Order{}, fmt.Errorf("failed to save order: %w", err)
	}

	log.Printf("[CreateOrder] Order saved successfully with ID: %d", createdOrder.ID)

	return createdOrder, nil
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
