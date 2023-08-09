package repository

import (
	"echoinventory/models"
	"echoinventory/schemas"
	"errors"

	"gorm.io/gorm"
)

type repositoryCategory struct {
	db *gorm.DB
}

func NewRepositoryCategory(db *gorm.DB) *repositoryCategory {
	return &repositoryCategory{
		db: db,
	}
}

func (r *repositoryCategory) EntityResults() (*[]models.ModelCategory, error) {
	var category []models.ModelCategory

	db := r.db.Model(&category)

	checkCategory := db.Debug().Find(&category)

	if checkCategory.RowsAffected < 1 {
		return &category, errors.New("error category")
	}

	return &category, nil
}

func (r *repositoryCategory) EntityCreate(input *schemas.SchemaCategory) (*models.ModelCategory, error) {
	var category models.ModelCategory

	category.Name = input.Name

	db := r.db.Model(&category)

	checkCategoryName := db.Debug().First(&category, "name=?", input.Name)

	if checkCategoryName.RowsAffected > 0 {
		return &category, checkCategoryName.Error
	}

	addCategory := db.Debug().Create(&category).Commit()

	if addCategory.RowsAffected < 1 {
		return &category, addCategory.Error
	}

	return &category, nil

}

func (r *repositoryCategory) EntityResult(input *schemas.SchemaCategory) (*models.ModelCategory, error) {
	var category models.ModelCategory
	category.ID = input.ID

	db := r.db.Model(&category)

	checkCategory := db.Debug().First(&category, "id=?", input.ID)

	if checkCategory.RowsAffected < 1 {
		return &category, checkCategory.Error
	}

	return &category, nil
}

func (r *repositoryCategory) EntityUpdate(input *schemas.SchemaCategory) (*models.ModelCategory, error) {
	var category models.ModelCategory

	category.Name = input.Name

	db := r.db.Model(&category)

	checkCategory := db.Debug().First(&category, "id=?", input.ID)

	if checkCategory.RowsAffected < 1 {
		return &category, checkCategory.Error
	}

	updateCategory := db.Debug().Save(&category).Commit()

	if updateCategory.RowsAffected < 1 {
		return &category, updateCategory.Error
	}

	return &category, nil
}

func (r *repositoryCategory) EntityDelete(input *schemas.SchemaCategory) (*models.ModelCategory, error) {
	var category models.ModelCategory

	db := r.db.Model(&category)

	checkCategory := db.Debug().First(&category, "id=?", input.ID)

	if checkCategory.RowsAffected < 1 {
		return &category, checkCategory.Error
	}

	deleteCategory := db.Debug().Delete(&category).Commit()

	if deleteCategory.RowsAffected < 1 {
		return &category, deleteCategory.Error
	}

	return &category, nil
}
