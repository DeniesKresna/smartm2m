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
	Name         string  `json:"name" validate:"required" valerr:"nama harus diisi"`
	Price        float64 `json:"price" validate:"required,number,gt=0" valerr:"Harga harus diisi, berupa angka, dan lebih dari 0"`
	Availability *int    `json:"availability" validate:"required,numeric,gte=0" valerr:"Availibility harus diisi, berupa angka dan tidak boleh negatif"`
	IsActive     *bool   `json:"is_active" validate:"required" valerr:"isActive harus diisi"`
}
