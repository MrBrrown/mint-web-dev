package productrepo

import "time"

type Product struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"desc"`
	Price       float64        `json:"price"`
	Attributes  map[string]any `json:"attribs"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"update_at"`
}
