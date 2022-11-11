package repository

import (
	"echoinventory/models"
	"echoinventory/schemas"

	"gorm.io/gorm"
)

type repositorySale struct {
	db *gorm.DB
}

func NewRepositorySale(db *gorm.DB) *repositorySale {
	return &repositorySale{
		db: db,
	}
}

func (r *repositorySale) EntityResults() (*[]models.ModelSale, error) {
	var sale []models.ModelSale

	db := r.db.Model(&sale)

	checkSale := db.Debug().Find(&sale)

	if checkSale.RowsAffected < 1 {
		return &sale, checkSale.Error
	}

	return &sale, nil
}

func (r *repositorySale) EntityCreate(input *schemas.SchemaSale) (*models.ModelSale, error) {
	var sale models.ModelSale

	sale.Name = input.Name
	sale.Telepon = input.Telepon
	sale.Email = input.Email
	sale.Alamat = input.Alamat

	db := r.db.Model(&sale)

	checkSaleName := db.Debug().First(&sale, "name=?", input.Name)

	if checkSaleName.RowsAffected > 0 {
		return &sale, checkSaleName.Error
	}

	addSale := db.Debug().Create(&sale).Commit()

	if addSale.RowsAffected < 1 {
		return &sale, addSale.Error
	}

	return &sale, nil

}

func (r *repositorySale) EntityResult(input *schemas.SchemaSale) (*models.ModelSale, error) {
	var sale models.ModelSale
	sale.ID = input.ID

	db := r.db.Model(&sale)

	checkSale := db.Debug().First(&sale, "id=?", input.ID)

	if checkSale.RowsAffected < 1 {
		return &sale, checkSale.Error
	}

	return &sale, nil
}

func (r *repositorySale) EntityUpdate(input *schemas.SchemaSale) (*models.ModelSale, error) {
	var sale models.ModelSale

	sale.Name = input.Name
	sale.Telepon = input.Telepon
	sale.Email = input.Email
	sale.Alamat = input.Alamat

	db := r.db.Model(&sale)

	checkSale := db.Debug().First(&sale, "id=?", input.ID)

	if checkSale.RowsAffected < 1 {
		return &sale, checkSale.Error
	}

	updateSale := db.Debug().Save(&sale).Commit()

	if updateSale.RowsAffected < 1 {
		return &sale, updateSale.Error
	}

	return &sale, nil
}

func (r *repositorySale) EntityDelete(input *schemas.SchemaSale) (*models.ModelSale, error) {
	var sale models.ModelSale

	db := r.db.Model(&sale)

	checkSale := db.Debug().First(&sale, "id=?", input.ID)

	if checkSale.RowsAffected < 1 {
		return &sale, checkSale.Error
	}

	deleteSale := db.Debug().Delete(&sale).Commit()

	if deleteSale.RowsAffected < 1 {
		return &sale, deleteSale.Error
	}

	return &sale, nil
}
