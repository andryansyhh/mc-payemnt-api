package usecase

import (
	"log"
	"mc-payment-api/models"
)

type BookUsecase interface {
	PostBook(req models.BookRequest) error
	GetAllBooks() (res []models.BookResponse, err error)
	GetBookByID(ID int) (res models.BookResponse, err error)
	UpdateBookByID(ID int, req models.BookRequest) error
	DeleteBookByID(ID int) error
}

func (b McPaymentUsecase) PostBook(req models.BookRequest) error {
	category, err := b.repo.GetCategoryByID(req.KategoriID)
	if category.ID == 0 {
		return models.ErrCategoryIDNotFound
	}
	if err != nil {
		return err
	}

	author, err := b.repo.GetAuthorByID(req.PenulisID)
	if author.ID == 0 {
		return models.ErrAuthorIDNotFound
	}
	if err != nil {
		return err
	}


	err = b.repo.PostBookData(req)
	if err != nil {
		return err
	}
	return nil
}

func (b McPaymentUsecase) GetAllBooks() (res []models.BookResponse, err error) {
	res, err = b.repo.GetAllBooks()
	log.Println(res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (b McPaymentUsecase) GetBookByID(ID int) (res models.BookResponse, err error) {
	res, err = b.repo.GetBookByID(ID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (b McPaymentUsecase) UpdateBookByID(ID int, req models.BookRequest) error {
	res, err := b.repo.GetBookByID(ID)
	if err != nil {
		return err
	}

	if res.ID == 0 {
		return models.ErrIdNotFound
	}

	if req.Judul == "" {
		req.Judul = res.Judul
	}

	if req.TahunTerbit == "" {
		req.TahunTerbit = res.TahunTerbit
	}

	if req.JumlahHalaman == 0 {
		req.JumlahHalaman = res.JumlahHalaman
	}

	if req.KategoriID == 0 {
		req.KategoriID = res.KategoriID
	}

	if req.PenulisID == 0 {
		req.PenulisID = res.PenulisID
	}

	err = b.repo.UpdateBookByID(ID, req)
	if err != nil {
		return err
	}

	return nil

}

func (b McPaymentUsecase) DeleteBookByID(ID int) error {
	res, err := b.repo.GetBookByID(ID)
	if err != nil {
		return err
	}

	if res.ID == 0 {
		return models.ErrIdNotFound
	}

	err = b.repo.DeleteBookByID(ID)
	if err != nil {
		return err
	}
	return nil
}
