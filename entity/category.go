package entity

import (
	"echoinventory/models"
	"echoinventory/schemas"
)

type EntityCategory interface {
	EntityResults() (*[]models.ModelCategory, error)
	EntityCreate(input *schemas.SchemaCategory) (*models.ModelCategory, error)
	EntityResult(input *schemas.SchemaCategory) (*models.ModelCategory, error)
	EntityUpdate(input *schemas.SchemaCategory) (*models.ModelCategory, error)
	EntityDelete(input *schemas.SchemaCategory) (*models.ModelCategory, error)
}
