package models

import "time"

type Review struct {
	ID          uint64    `json:"id"`
	ProductID   uint64    `json:"product_id"`
	Message     string    `json:"message"`
	CreatedTime time.Time `json:"created_time"`
}
