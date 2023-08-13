package entity

import (
	"echoinventory/models"
	"echoinventory/schemas"
)

type EntityAuth interface {
	EntityRegister(input *schemas.SchemaUser) (*models.ModelUser, error)
	EntityLogin(input *schemas.SchemaUser) (*models.ModelUser, error)
}

