package usecase

import (
	"log"
	"mc-payment-api/models"
)

type AuthorUsecase interface {
	PostAuthor(req models.AuthorRequest) error
	GetAllAuthors() (res []models.AuthorResponse, err error)
	GetAuthorByID(ID int) (res models.AuthorResponse, err error)
	UpdateAuthorByID(ID int, req models.AuthorRequest) error
	DeleteAuthorByID(ID int) error
}

func (b McPaymentUsecase) PostAuthor(req models.AuthorRequest) error {
	err := b.repo.PostAuthorData(req)
	if err != nil {
		return err
	}
	return nil
}

func (b McPaymentUsecase) GetAllAuthors() (res []models.AuthorResponse, err error) {
	res, err = b.repo.GetAllAuthors()
	log.Println(res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (b McPaymentUsecase) GetAuthorByID(ID int) (res models.AuthorResponse, err error) {
	res, err = b.repo.GetAuthorByID(ID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (b McPaymentUsecase) UpdateAuthorByID(ID int, req models.AuthorRequest) error {
	res, err := b.repo.GetAuthorByID(ID)
	if err != nil {
		return err
	}

	if res.ID == 0 {
		return models.ErrIdNotFound
	}

	if req.Nama == "" {
		req.Nama = res.Nama
	}

	if req.TangalLahir == "" {
		req.TangalLahir = res.TangalLahir
	}

	err = b.repo.UpdateAuthorByID(ID, req)
	if err != nil {
		return err
	}

	return nil

}

func (b McPaymentUsecase) DeleteAuthorByID(ID int) error {
	res, err := b.repo.GetAuthorByID(ID)
	if err != nil {
		return err
	}

	if res.ID == 0 {
		return models.ErrIdNotFound
	}

	err = b.repo.DeleteAuthorByID(ID)
	if err != nil {
		return err
	}
	return nil
}
