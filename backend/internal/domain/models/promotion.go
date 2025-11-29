package models

import "time"

type Promotion struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	DiscountPercent float64   `json:"discount_percent"`
	StartDate       string    `json:"start_date"`
	EndDate         string    `json:"end_date"`
	Active          bool      `json:"active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}