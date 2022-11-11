package entity

import (
	"echoinventory/models"
	"echoinventory/schemas"
)

type EntityCustomer interface {
	EntityResults() (*[]models.ModelCustomer, error)
	EntityCreate(input *schemas.SchemaCustomer) (*models.ModelCustomer, error)
	EntityResult(input *schemas.SchemaCustomer) (*models.ModelCustomer, error)
	EntityUpdate(input *schemas.SchemaCustomer) (*models.ModelCustomer, error)
	EntityDelete(input *schemas.SchemaCustomer) (*models.ModelCustomer, error)
}
