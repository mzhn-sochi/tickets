package entity

import (
	"time"
)

type Ticket struct {
	Id          string     `json:"id" db:"id"`
	UserId      string     `json:"user_id" db:"user_id"`
	ShopName    string     `json:"shop_name" db:"shop_name"`
	ShopAddress string     `json:"shop_id" db:"shop_address"`
	ImageUrl    string     `json:"image_url" db:"image_url"`
	Status      string     `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
	Reason      *string    `json:"reason" db:"reason"`
}
