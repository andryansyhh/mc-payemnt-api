package repository

import (
	"mc-payment-api/models"
	query "mc-payment-api/repository/query"
)

type AuthorRepo interface {
	PostAuthorData(input models.AuthorRequest) error
	GetAllAuthors() (res []models.AuthorResponse, err error)
	GetAuthorByID(ID int) (res models.AuthorResponse, err error)
	DeleteAuthorByID(ID int) error
	UpdateAuthorByID(ID int, req models.AuthorRequest) error
}

func (b *Repo) PostAuthorData(req models.AuthorRequest) error {
	_, err := b.db.Exec(
		query.PostAuthor,
		req.Nama,
		req.TangalLahir,
	)

	if err != nil {
		return err
	}

	return nil
}

func (b *Repo) GetAllAuthors() (res []models.AuthorResponse, err error) {
	var qry = query.GetAllAuthors

	row, err := b.db.Query(
		qry,
	)
	if err != nil {
		return res, err
	}

	for row.Next() {
		temp := models.AuthorResponse{}
		err = row.Scan(
			&temp.ID,
			&temp.Nama,
			&temp.TangalLahir,
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

func (b *Repo) GetAuthorByID(ID int) (res models.AuthorResponse, err error) {
	err = b.db.QueryRow(
		query.GetAuthorByID,
		ID,
	).Scan(
		&res.ID,
		&res.Nama,
		&res.TangalLahir,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (b *Repo) UpdateAuthorByID(ID int, req models.AuthorRequest) error {
	_, err := b.db.Exec(
		query.UpdateAuthorByID,
		ID,
		req.Nama,
		req.TangalLahir,
	)

	if err != nil {
		return err
	}
	return nil
}

func (b *Repo) DeleteAuthorByID(ID int) error {
	_, err := b.db.Exec(
		query.DeleteAuthorByID,
		ID,
	)

	if err != nil {
		return err
	}
	return nil
}
