package entity

import (
	"echoinventory/models"
	"echoinventory/schemas"
)

type EntityProductMasuk interface {
	EntityResults() (*[]models.ModelProductMasuk, error)
	EntityCreate(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, error)
	EntityResult(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, error)
	EntityUpdate(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, error)
	EntityDelete(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, error)
}
