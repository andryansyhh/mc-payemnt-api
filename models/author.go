package models

import "time"

type AuthorRequest struct {
	Nama        string `json:"nama"`
	TangalLahir string `json:"tanggal_lahir"`
}

type AuthorResponse struct {
	ID          int       `json:"id"`
	Nama        string    `json:"nama"`
	TangalLahir string    `json:"tanggal_lahir"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
