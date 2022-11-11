package services

import (
	"echoinventory/entity"
	"echoinventory/models"
	"echoinventory/schemas"
)

type serviceProduct struct {
	product entity.EntityProduct
}

func NewServiceProduct(product entity.EntityProduct) *serviceProduct {
	return &serviceProduct{
		product: product,
	}
}

func (s *serviceProduct) EntityResults() (*[]models.ModelProduct, error) {
	res, err := s.product.EntityResults()

	return res, err
}

func (s *serviceProduct) EntityCreate(input *schemas.SchemaProduct) (*models.ModelProduct, error) {
	var product schemas.SchemaProduct

	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	res, err := s.product.EntityCreate(&product)

	return res, err
}

func (s *serviceProduct) EntityResult(input *schemas.SchemaProduct) (*models.ModelProduct, error) {
	var product schemas.SchemaProduct
	product.ID = input.ID

	res, err := s.product.EntityResult(&product)

	return res, err
}

func (s *serviceProduct) EntityUpdate(input *schemas.SchemaProduct) (*models.ModelProduct, error) {
	var product schemas.SchemaProduct

	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	res, err := s.product.EntityUpdate(&product)

	return res, err
}

func (s *serviceProduct) EntityDelete(input *schemas.SchemaProduct) (*models.ModelProduct, error) {
	var product schemas.SchemaProduct
	product.ID = input.ID

	res, err := s.product.EntityDelete(&product)

	return res, err
}
