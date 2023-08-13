package repository

import (
	"echoinventory/models"
	"echoinventory/schemas"

	"gorm.io/gorm"
)

type repositoryProduct struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repositoryProduct {
	return &repositoryProduct{
		db: db,
	}
}

func (r *repositoryProduct) EntityResults() (*[]models.ModelProduct, error) {
	var product []models.ModelProduct

	db := r.db.Model(&product)

	checkProduct := db.Debug().Preload("Category").Find(&product)

	if checkProduct.RowsAffected < 1 {
		return &product, checkProduct.Error
	}

	return &product, nil
}

func (r *repositoryProduct) EntityCreate(input *schemas.SchemaProduct) (*models.ModelProduct, error) {
	var product models.ModelProduct

	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	db := r.db.Model(&product)

	checkProductName := db.Debug().First(&product, "name=?", input.Name)

	if checkProductName.RowsAffected > 0 {
		return &product, checkProductName.Error
	}

	addProduct := db.Debug().Create(&product).Commit()

	if addProduct.RowsAffected < 1 {
		return &product, addProduct.Error
	}

	return &product, nil

}

func (r *repositoryProduct) EntityResult(input *schemas.SchemaProduct) (*models.ModelProduct, error) {
	var product models.ModelProduct
	product.ID = input.ID

	db := r.db.Model(&product)

	checkProduct := db.Debug().First(&product, "id=?", input.ID)

	if checkProduct.RowsAffected < 1 {
		return &product, checkProduct.Error
	}

	return &product, nil
}

func (r *repositoryProduct) EntityUpdate(input *schemas.SchemaProduct) (*models.ModelProduct, error) {
	var product models.ModelProduct

	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	db := r.db.Model(&product)

	checkProduct := db.Debug().First(&product, "id=?", input.ID)

	if checkProduct.RowsAffected < 1 {
		return &product, checkProduct.Error
	}

	updateProduct := db.Debug().Save(&product).Commit()

	if updateProduct.RowsAffected < 1 {
		return &product, updateProduct.Error
	}

	return &product, nil
}

func (r *repositoryProduct) EntityDelete(input *schemas.SchemaProduct) (*models.ModelProduct, error) {
	var product models.ModelProduct

	db := r.db.Model(&product)

	checkProduct := db.Debug().First(&product, "id=?", input.ID)

	if checkProduct.RowsAffected < 1 {
		return &product, checkProduct.Error
	}

	deleteProduct := db.Debug().Delete(&product).Commit()

	if deleteProduct.RowsAffected < 1 {
		return &product, deleteProduct.Error
	}

	return &product, nil
}
