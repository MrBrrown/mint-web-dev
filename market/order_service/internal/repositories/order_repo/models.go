package orderrepo

import "time"

type Order struct {
	ID         uint             `json:"id"`
	TotalPrice float64          `json:"total_price"`
	Items      []map[string]any `json:"items"`
	UserInfo   map[string]any   `json:"user_info"`
	Status     string           `json:"status"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
}
