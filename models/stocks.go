package models

import "time"

type Stock struct {
	ID           int64     `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Price        float64   `json:"price" db:"price"`
	Availability int       `json:"availability" db:"availability"`
	IsActive     *bool     `json:"is_active" db:"is_active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type StockCreateRequest struct {
	Name         string  `json:"name" db:"name"`
	Price        float64 `json:"price" db:"price"`
	Availability int     `json:"availability"`
	IsActive     *bool   `json:"is_active"`
}
