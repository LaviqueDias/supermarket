package models

import "time"

type Cart struct {
	ID        int       `json:"id"`
	ClientID  int       `json:"client_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}