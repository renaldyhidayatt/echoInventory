package repository

import (
	"echoinventory/models"
	"echoinventory/schemas"

	"gorm.io/gorm"
)

type repositoryProductKeluar struct {
	db *gorm.DB
}

func NewRepositoryProductKeluar(db *gorm.DB) *repositoryProductKeluar {
	return &repositoryProductKeluar{
		db: db,
	}
}

func (r *repositoryProductKeluar) EntityResults() (*[]models.ModelProductKeluar, error) {
	var productkeluar []models.ModelProductKeluar

	db := r.db.Model(&productkeluar)

	checkProductKeluar := db.Debug().Find(&productkeluar)

	if checkProductKeluar.RowsAffected < 1 {
		return &productkeluar, checkProductKeluar.Error
	}

	return &productkeluar, nil
}

func (r *repositoryProductKeluar) EntityCreate(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, error) {
	var productkeluar models.ModelProductKeluar

	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	db := r.db.Model(&productkeluar)

	addProductKeluar := db.Debug().Create(&productkeluar).Commit()

	if addProductKeluar.RowsAffected < 1 {
		return &productkeluar, addProductKeluar.Error
	}

	return &productkeluar, nil

}

func (r *repositoryProductKeluar) EntityResult(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, error) {
	var productkeluar models.ModelProductKeluar
	productkeluar.ID = input.ID

	db := r.db.Model(&productkeluar)

	checkProductKeluar := db.Debug().First(&productkeluar, "id=?", input.ID)

	if checkProductKeluar.RowsAffected < 1 {
		return &productkeluar, checkProductKeluar.Error
	}

	return &productkeluar, nil
}

func (r *repositoryProductKeluar) EntityUpdate(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, error) {
	var productkeluar models.ModelProductKeluar

	productkeluar.ID = input.ID
	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	db := r.db.Model(&productkeluar)

	checkProductKeluar := db.Debug().First(&productkeluar, "id=?", input.ID)

	if checkProductKeluar.RowsAffected < 1 {
		return &productkeluar, checkProductKeluar.Error
	}

	updateProductKeluar := db.Debug().Save(&productkeluar).Commit()

	if updateProductKeluar.RowsAffected < 1 {
		return &productkeluar, updateProductKeluar.Error
	}

	return &productkeluar, nil
}

func (r *repositoryProductKeluar) EntityDelete(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, error) {
	var productkeluar models.ModelProductKeluar

	db := r.db.Model(&productkeluar)

	checkProductKeluar := db.Debug().First(&productkeluar, "id=?", input.ID)

	if checkProductKeluar.RowsAffected < 1 {
		return &productkeluar, checkProductKeluar.Error
	}

	deleteProductKeluar := db.Debug().Delete(&productkeluar).Commit()

	if deleteProductKeluar.RowsAffected < 1 {
		return &productkeluar, deleteProductKeluar.Error
	}

	return &productkeluar, nil
}
