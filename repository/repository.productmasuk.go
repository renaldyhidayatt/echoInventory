package repository

import (
	"echoinventory/models"
	"echoinventory/schemas"

	"gorm.io/gorm"
)

type repositoryProductMasuk struct {
	db *gorm.DB
}

func NewRepositoryProductMasuk(db *gorm.DB) *repositoryProductMasuk {
	return &repositoryProductMasuk{
		db: db,
	}
}

func (r *repositoryProductMasuk) EntityResults() (*[]models.ModelProductMasuk, error) {
	var productmasuk []models.ModelProductMasuk

	db := r.db.Model(&productmasuk)

	checkProductMasuk := db.Debug().Find(&productmasuk)

	if checkProductMasuk.RowsAffected < 1 {
		return &productmasuk, checkProductMasuk.Error
	}

	return &productmasuk, nil
}

func (r *repositoryProductMasuk) EntityCreate(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, error) {
	var productmasuk models.ModelProductMasuk

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	db := r.db.Model(&productmasuk)

	checkProductMasuk := db.Debug().First(&productmasuk, "name=?", input.Name)

	if checkProductMasuk.RowsAffected > 0 {
		return &productmasuk, checkProductMasuk.Error
	}

	addProductMasuk := db.Debug().Create(&productmasuk).Commit()

	if addProductMasuk.RowsAffected < 1 {
		return &productmasuk, addProductMasuk.Error
	}

	return &productmasuk, nil

}

func (r *repositoryProductMasuk) EntityResult(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, error) {
	var productmasuk models.ModelProductMasuk
	productmasuk.ID = input.ID

	db := r.db.Model(&productmasuk)

	checkProductMasuk := db.Debug().First(&productmasuk, "id=?", input.ID)

	if checkProductMasuk.RowsAffected < 1 {
		return &productmasuk, checkProductMasuk.Error
	}

	return &productmasuk, nil
}

func (r *repositoryProductMasuk) EntityUpdate(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, error) {
	var productmasuk models.ModelProductMasuk

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	db := r.db.Model(&productmasuk)

	checkProductMasuk := db.Debug().First(&productmasuk, "id=?", input.ID)

	if checkProductMasuk.RowsAffected < 1 {
		return &productmasuk, checkProductMasuk.Error
	}

	updateProductMasuk := db.Debug().Save(&productmasuk).Commit()

	if updateProductMasuk.RowsAffected < 1 {
		return &productmasuk, updateProductMasuk.Error
	}

	return &productmasuk, nil
}

func (r *repositoryProductMasuk) EntityDelete(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, error) {
	var productmasuk models.ModelProductMasuk

	db := r.db.Model(&productmasuk)

	checkProductMasuk := db.Debug().First(&productmasuk, "id=?", input.ID)

	if checkProductMasuk.RowsAffected < 1 {
		return &productmasuk, checkProductMasuk.Error
	}

	deleteProductMasuk := db.Debug().Delete(&productmasuk).Commit()

	if deleteProductMasuk.RowsAffected < 1 {
		return &productmasuk, deleteProductMasuk.Error
	}

	return &productmasuk, nil
}
