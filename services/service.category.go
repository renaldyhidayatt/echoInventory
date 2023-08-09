package services

import (
	"echoinventory/entity"
	"echoinventory/models"
	"echoinventory/schemas"
	"fmt"
)

type serviceCategory struct {
	category entity.EntityCategory
}

func NewServiceCategory(category entity.EntityCategory) *serviceCategory {
	return &serviceCategory{
		category: category,
	}
}

func (s *serviceCategory) EntityCreate(input *schemas.SchemaCategory) (*models.ModelCategory, error) {
	var category schemas.SchemaCategory

	category.Name = input.Name

	res, err := s.category.EntityCreate(&category)

	return res, err
}

func (s *serviceCategory) EntityResults() (*[]models.ModelCategory, error) {
	res, err := s.category.EntityResults()

	if err != nil {
		return nil, fmt.Errorf("failed error %w", err)
	}

	return res, err
}

func (s *serviceCategory) EntityResult(input *schemas.SchemaCategory) (*models.ModelCategory, error) {
	var category schemas.SchemaCategory
	category.ID = input.ID

	res, err := s.category.EntityResult(&category)

	return res, err
}

func (s *serviceCategory) EntityUpdate(input *schemas.SchemaCategory) (*models.ModelCategory, error) {
	var category schemas.SchemaCategory

	category.Name = input.Name

	res, err := s.category.EntityUpdate(&category)

	return res, err
}

func (s *serviceCategory) EntityDelete(input *schemas.SchemaCategory) (*models.ModelCategory, error) {
	var category schemas.SchemaCategory
	category.ID = input.ID

	res, err := s.category.EntityDelete(&category)

	return res, err
}
