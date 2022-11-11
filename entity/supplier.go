package entity

import (
	"echoinventory/models"
	"echoinventory/schemas"
)

type EntitySupplier interface {
	EntityResults() (*[]models.ModelSupplier, error)
	EntityCreate(input *schemas.SchemaSupplier) (*models.ModelSupplier, error)
	EntityResult(input *schemas.SchemaSupplier) (*models.ModelSupplier, error)
	EntityUpdate(input *schemas.SchemaSupplier) (*models.ModelSupplier, error)
	EntityDelete(input *schemas.SchemaSupplier) (*models.ModelSupplier, error)
}
