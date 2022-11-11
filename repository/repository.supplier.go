package repository

import (
	"echoinventory/models"
	"echoinventory/schemas"

	"gorm.io/gorm"
)

type repositorySupplier struct {
	db *gorm.DB
}

func NewRepositorySupplier(db *gorm.DB) *repositorySupplier {
	return &repositorySupplier{
		db: db,
	}
}

func (r *repositorySupplier) EntityResults() (*[]models.ModelSupplier, error) {
	var supplier []models.ModelSupplier

	db := r.db.Model(&supplier)

	checkSupplier := db.Debug().Find(&supplier)

	if checkSupplier.RowsAffected < 1 {
		return &supplier, checkSupplier.Error
	}

	return &supplier, nil
}

func (r *repositorySupplier) EntityCreate(input *schemas.SchemaSupplier) (*models.ModelSupplier, error) {
	var supplier models.ModelSupplier

	supplier.Name = input.Name
	supplier.Telepon = input.Telepon
	supplier.Email = input.Email
	supplier.Alamat = input.Alamat

	db := r.db.Model(&supplier)

	checkSupplier := db.Debug().First(&supplier, "name=?", input.Name)

	if checkSupplier.RowsAffected > 0 {
		return &supplier, checkSupplier.Error
	}

	addSupplier := db.Debug().Create(&supplier).Commit()

	if addSupplier.RowsAffected < 1 {
		return &supplier, addSupplier.Error
	}

	return &supplier, nil

}

func (r *repositorySupplier) EntityResult(input *schemas.SchemaSupplier) (*models.ModelSupplier, error) {
	var supplier models.ModelSupplier
	supplier.ID = input.ID

	db := r.db.Model(&supplier)

	checkSupplier := db.Debug().First(&supplier, "id=?", input.ID)

	if checkSupplier.RowsAffected < 1 {
		return &supplier, checkSupplier.Error
	}

	return &supplier, nil
}

func (r *repositorySupplier) EntityUpdate(input *schemas.SchemaSupplier) (*models.ModelSupplier, error) {
	var supplier models.ModelSupplier

	supplier.Name = input.Name
	supplier.Telepon = input.Telepon
	supplier.Email = input.Email
	supplier.Alamat = input.Alamat

	db := r.db.Model(&supplier)

	checkSupplier := db.Debug().First(&supplier, "id=?", input.ID)

	if checkSupplier.RowsAffected < 1 {
		return &supplier, checkSupplier.Error
	}

	updateSupplier := db.Debug().Save(&supplier).Commit()

	if updateSupplier.RowsAffected < 1 {
		return &supplier, updateSupplier.Error
	}

	return &supplier, nil
}

func (r *repositorySupplier) EntityDelete(input *schemas.SchemaSupplier) (*models.ModelSupplier, error) {
	var supplier models.ModelSupplier

	db := r.db.Model(&supplier)

	checkSupplier := db.Debug().First(&supplier, "id=?", input.ID)

	if checkSupplier.RowsAffected < 1 {
		return &supplier, checkSupplier.Error
	}

	deleteSupplier := db.Debug().Delete(&supplier).Commit()

	if deleteSupplier.RowsAffected < 1 {
		return &supplier, deleteSupplier.Error
	}

	return &supplier, nil
}
