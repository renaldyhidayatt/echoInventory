package entity

import (
	"echoinventory/models"
	"echoinventory/schemas"
)

type EntitySale interface {
	EntityResults() (*[]models.ModelSale, error)
	EntityCreate(input *schemas.SchemaSale) (*models.ModelSale, error)
	EntityResult(input *schemas.SchemaSale) (*models.ModelSale, error)
	EntityUpdate(input *schemas.SchemaSale) (*models.ModelSale, error)
	EntityDelete(input *schemas.SchemaSale) (*models.ModelSale, error)
}
