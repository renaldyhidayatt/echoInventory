package repository

import (
	"echoinventory/models"
	"echoinventory/schemas"

	"gorm.io/gorm"
)

type repositoryCustomer struct {
	db *gorm.DB
}

func NewRepositoryCustomer(db *gorm.DB) *repositoryCustomer {
	return &repositoryCustomer{
		db: db,
	}
}

func (r *repositoryCustomer) EntityResults() (*[]models.ModelCustomer, error) {
	var customer []models.ModelCustomer

	db := r.db.Model(&customer)

	checkCustomer := db.Debug().Find(&customer)

	if checkCustomer.RowsAffected < 1 {
		return &customer, checkCustomer.Error
	}

	return &customer, nil
}

func (r *repositoryCustomer) EntityCreate(input *schemas.SchemaCustomer) (*models.ModelCustomer, error) {
	var customer models.ModelCustomer

	customer.Name = input.Name
	customer.Telepon = input.Telepon
	customer.Email = input.Email
	customer.Alamat = input.Alamat

	db := r.db.Model(&customer)

	checkCustomerName := db.Debug().First(&customer, "name=?", input.Name)

	if checkCustomerName.RowsAffected > 0 {
		return &customer, checkCustomerName.Error
	}

	addCustomer := db.Debug().Create(&customer).Commit()

	if addCustomer.RowsAffected < 1 {
		return &customer, addCustomer.Error
	}

	return &customer, nil

}

func (r *repositoryCustomer) EntityResult(input *schemas.SchemaCustomer) (*models.ModelCustomer, error) {
	var customer models.ModelCustomer
	customer.ID = input.ID

	db := r.db.Model(&customer)

	checkCustomer := db.Debug().First(&customer, "id=?", input.ID)

	if checkCustomer.RowsAffected < 1 {
		return &customer, checkCustomer.Error
	}

	return &customer, nil
}

func (r *repositoryCustomer) EntityUpdate(input *schemas.SchemaCustomer) (*models.ModelCustomer, error) {
	var customer models.ModelCustomer

	customer.Name = input.Name
	customer.Telepon = input.Telepon
	customer.Email = input.Email
	customer.Alamat = input.Alamat

	db := r.db.Model(&customer)

	checkCustomer := db.Debug().First(&customer, "id=?", input.ID)

	if checkCustomer.RowsAffected < 1 {
		return &customer, checkCustomer.Error
	}

	updateCustomer := db.Debug().Save(&customer).Commit()

	if updateCustomer.RowsAffected < 1 {
		return &customer, updateCustomer.Error
	}

	return &customer, nil
}

func (r *repositoryCustomer) EntityDelete(input *schemas.SchemaCustomer) (*models.ModelCustomer, error) {
	var customer models.ModelCustomer

	db := r.db.Model(&customer)

	checkCustomer := db.Debug().First(&customer, "id=?", input.ID)

	if checkCustomer.RowsAffected < 1 {
		return &customer, checkCustomer.Error
	}

	deleteCustomer := db.Debug().Delete(&customer).Commit()

	if deleteCustomer.RowsAffected < 1 {
		return &customer, deleteCustomer.Error
	}

	return &customer, nil
}
