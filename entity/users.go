package entity

import (
	"echoinventory/models"
	"echoinventory/schemas"
)

type EntityUsers interface {
	EntityResults() (*[]models.ModelUser, error)
	EntityCreate(input *schemas.SchemaUser) (*models.ModelUser, error)
	EntityResult(input *schemas.SchemaUser) (*models.ModelUser, error)
	EntityUpdate(input *schemas.SchemaUser) (*models.ModelUser, error)
	EntityDelete(input *schemas.SchemaUser) (*models.ModelUser, error)
}
