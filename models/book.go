package models

import "time"

type BookRequest struct {
	Judul         string `json:"judul"`
	TahunTerbit   string `json:"tahun_terbit"`
	JumlahHalaman int    `json:"jumlah_halaman"`
	KategoriID    int    `json:"kategori_id"`
	PenulisID     int    `json:"penulis_id"`
}

type BookResponse struct {
	ID            int       `json:"id"`
	Judul         string    `json:"judul"`
	TahunTerbit   string    `json:"tahun_terbit"`
	JumlahHalaman int       `json:"jumlah_halaman"`
	KategoriID    int       `json:"kategori_id"`
	PenulisID     int       `json:"penulis_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}