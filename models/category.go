package models

import "time"

type CategoryRequest struct {
	Nama string `json:"nama"`
}

type CategoryResponse struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
