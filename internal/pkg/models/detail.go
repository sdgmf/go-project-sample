package models

import "time"

type Detail struct {
	ID    uint64 `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	CreatedTime time.Time `json:"created_time"`
}
