package services

import (
	"echoinventory/entity"
	"echoinventory/models"
	"echoinventory/schemas"
)

type serviceAuth struct {
	auth entity.EntityAuth
}

func NewServiceAuth(auth entity.EntityAuth) *serviceAuth {
	return &serviceAuth{auth: auth}
}

func (s *serviceAuth) EntityRegister(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var schema schemas.SchemaUser

	schema.FirstName = input.FirstName
	schema.LastName = input.LastName
	schema.Email = input.Email
	schema.Password = input.Password
	schema.Role = input.Role

	res, err := s.auth.EntityRegister(&schema)

	return res, err

}

func (s *serviceAuth) EntityLogin(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var schema schemas.SchemaUser

	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.auth.EntityLogin(&schema)

	return res, err
}
