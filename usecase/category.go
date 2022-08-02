package usecase

import (
	"mc-payment-api/models"
)

type CategoryUsecase interface {
	PostCategory(req models.CategoryRequest) error
	GetAllCategories() (res []models.CategoryResponse, err error)
	GetCategoryByID(ID int) (res models.CategoryResponse, err error)
	UpdateCategoryByID(ID int, req models.CategoryRequest) error
	DeleteCategoryByID(ID int) error
}

func (b McPaymentUsecase) PostCategory(req models.CategoryRequest) error {
	err := b.repo.PostCategoryData(req)
	if err != nil {
		return err
	}
	return nil
}

func (b McPaymentUsecase) GetAllCategories() (res []models.CategoryResponse, err error) {
	res, err = b.repo.GetAllCategories()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (b McPaymentUsecase) GetCategoryByID(ID int) (res models.CategoryResponse, err error) {
	res, err = b.repo.GetCategoryByID(ID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (b McPaymentUsecase) UpdateCategoryByID(ID int, req models.CategoryRequest) error {
	res, err := b.repo.GetCategoryByID(ID)
	if err != nil {
		return err
	}

	if res.ID == 0 {
		return models.ErrIdNotFound
	}

	if req.Nama == "" {
		req.Nama = res.Nama
	}

	err = b.repo.UpdateCategoryByID(ID, req)
	if err != nil {
		return err
	}

	return nil

}

func (b McPaymentUsecase) DeleteCategoryByID(ID int) error {
	res, err := b.repo.GetCategoryByID(ID)
	if err != nil {
		return err
	}

	if res.ID == 0 {
		return models.ErrIdNotFound
	}

	err = b.repo.DeleteCategoryByID(ID)
	if err != nil {
		return err
	}
	return nil
}
