package repository

import (
	"log"
	"mc-payment-api/models"
	query "mc-payment-api/repository/query"
)

type BookRepo interface {
	PostBookData(input models.BookRequest) error
	GetAllBooks() (res []models.BookResponse, err error)
	GetBookByID(ID int) (res models.BookResponse, err error)
	DeleteBookByID(ID int) error
	UpdateBookByID(ID int, req models.BookRequest) error
}

func (b *Repo) PostBookData(req models.BookRequest) error {
	_, err := b.db.Exec(
		query.PostBook,
		req.Judul,
		req.TahunTerbit,
		req.JumlahHalaman,
		req.KategoriID,
		req.PenulisID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (b *Repo) GetAllBooks() (res []models.BookResponse, err error) {
	var qry = query.GetAllBooks

	row, err := b.db.Query(
		qry,
	)
	if err != nil {
		return res, err
	}

	for row.Next() {
		temp := models.BookResponse{}
		err = row.Scan(
			&temp.ID,
			&temp.Judul,
			&temp.TahunTerbit,
			&temp.JumlahHalaman,
			&temp.KategoriID,
			&temp.PenulisID,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}
	return res, nil
}

func (b *Repo) GetBookByID(ID int) (res models.BookResponse, err error) {
	err = b.db.QueryRow(
		query.GetBookByID,
		ID,
	).Scan(
		&res.ID,
		&res.Judul,
		&res.TahunTerbit,
		&res.JumlahHalaman,
		&res.KategoriID,
		&res.PenulisID,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	log.Println(err)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (b *Repo) UpdateBookByID(ID int, req models.BookRequest) error {
	_, err := b.db.Exec(
		query.UpdateBookByID,
		ID,
		req.Judul,
		req.TahunTerbit,
		req.JumlahHalaman,
		req.KategoriID,
		req.PenulisID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (b *Repo) DeleteBookByID(ID int) error {
	_, err := b.db.Exec(
		query.DeleteBookByID,
		ID,
	)

	if err != nil {
		return err
	}
	return nil
}
