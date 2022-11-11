package services

import (
	"echoinventory/entity"
	"echoinventory/models"
	"echoinventory/schemas"
)

type serviceProductKeluar struct {
	productkeluar entity.EntityProductKeluar
}

func NewServiceProductKeluar(productkeluar entity.EntityProductKeluar) *serviceProductKeluar {
	return &serviceProductKeluar{
		productkeluar: productkeluar,
	}
}

func (s *serviceProductKeluar) EntityResults() (*[]models.ModelProductKeluar, error) {
	res, err := s.productkeluar.EntityResults()

	return res, err
}

func (s *serviceProductKeluar) EntityCreate(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, error) {
	var productkeluar schemas.SchemaProductKeluar

	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	res, err := s.productkeluar.EntityCreate(&productkeluar)

	return res, err
}

func (s *serviceProductKeluar) EntityResult(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, error) {
	var productkeluar schemas.SchemaProductKeluar
	productkeluar.ID = input.ID

	res, err := s.productkeluar.EntityResult(&productkeluar)

	return res, err
}

func (s *serviceProductKeluar) EntityUpdate(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, error) {
	var productkeluar schemas.SchemaProductKeluar

	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	res, err := s.productkeluar.EntityUpdate(&productkeluar)

	return res, err
}

func (s *serviceProductKeluar) EntityDelete(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, error) {
	var productkeluar schemas.SchemaProductKeluar
	productkeluar.ID = input.ID

	res, err := s.productkeluar.EntityDelete(&productkeluar)

	return res, err
}
