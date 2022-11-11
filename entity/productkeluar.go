package entity

import (
	"echoinventory/models"
	"echoinventory/schemas"
)

type EntityProductKeluar interface {
	EntityResults() (*[]models.ModelProductKeluar, error)
	EntityCreate(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, error)
	EntityResult(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, error)
	EntityUpdate(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, error)
	EntityDelete(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, error)
}
