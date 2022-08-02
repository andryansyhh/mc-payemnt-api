package repository

import (
	"log"
	"mc-payment-api/models"
	query "mc-payment-api/repository/query"
)

type CategoryRepo interface {
	PostCategoryData(input models.CategoryRequest) error
	GetAllCategories() (res []models.CategoryResponse, err error)
	GetCategoryByID(ID int) (res models.CategoryResponse, err error)
	DeleteCategoryByID(ID int) error
	UpdateCategoryByID(ID int, req models.CategoryRequest) error
}

func (b *Repo) PostCategoryData(req models.CategoryRequest) error {
	_, err := b.db.Exec(
		query.PostCategory,
		req.Nama,
	)

	if err != nil {
		return err
	}

	return nil
}

func (b *Repo) GetAllCategories() (res []models.CategoryResponse, err error) {
	var qry = query.GetAllCategories

	row, err := b.db.Query(
		qry,
	)
	if err != nil {
		return res, err
	}

	for row.Next() {
		temp := models.CategoryResponse{}
		err = row.Scan(
			&temp.ID,
			&temp.Nama,
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

func (b *Repo) GetCategoryByID(ID int) (res models.CategoryResponse, err error) {
	err = b.db.QueryRow(
		query.GetCategoryByID,
		ID,
	).Scan(
		&res.ID,
		&res.Nama,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	log.Println(err)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (b *Repo) UpdateCategoryByID(ID int, req models.CategoryRequest) error {
	_, err := b.db.Exec(
		query.UpdateCategoryByID,
		ID,
		req.Nama,
	)

	if err != nil {
		return err
	}
	return nil
}

func (b *Repo) DeleteCategoryByID(ID int) error {
	_, err := b.db.Exec(
		query.DeleteCategoryByID,
		ID,
	)

	if err != nil {
		return err
	}
	return nil
}
