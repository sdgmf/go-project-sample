package models

import "time"

type Rating struct {
	ID    uint64 `json:"id"`
	ProductID uint64 `json:"product_id"`
	Score uint32 `json:"score"`
	UpdatedTime time.Time `json:"updated_time"`
}
