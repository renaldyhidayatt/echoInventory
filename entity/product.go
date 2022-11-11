package entity

import (
	"echoinventory/models"
	"echoinventory/schemas"
)

type EntityProduct interface {
	EntityResults() (*[]models.ModelProduct, error)
	EntityCreate(input *schemas.SchemaProduct) (*models.ModelProduct, error)
	EntityResult(input *schemas.SchemaProduct) (*models.ModelProduct, error)
	EntityUpdate(input *schemas.SchemaProduct) (*models.ModelProduct, error)
	EntityDelete(input *schemas.SchemaProduct) (*models.ModelProduct, error)
}
