package entity

import (
	"time"
)

type Ticket struct {
	Id          string     `json:"id" db:"id"`
	UserId      string     `json:"user_id" db:"user_id"`
	ShopAddress string     `json:"shop_id" db:"shop_address"`
	ImageUrl    string     `json:"image_url" db:"image_url"`
	Status      string     `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}
