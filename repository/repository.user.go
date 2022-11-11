package repository

import (
	"echoinventory/models"
	"echoinventory/schemas"

	"gorm.io/gorm"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{
		db: db,
	}
}

func (r *repositoryUser) EntityResults() (*[]models.ModelUser, error) {
	var users []models.ModelUser

	db := r.db.Model(&users)

	checkUser := db.Debug().Find(&users)

	if checkUser.RowsAffected < 1 {
		return &users, checkUser.Error
	}

	return &users, nil
}

func (r *repositoryUser) EntityCreate(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user models.ModelUser
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Password = input.Password
	user.Role = input.Role

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user, "email=?", input.Email)

	if checkUser.RowsAffected > 0 {
		return &user, checkUser.Error
	}

	addUser := db.Debug().Create(&user).Commit()

	if addUser.RowsAffected < 1 {
		return &user, addUser.Error
	}

	return &user, nil

}

func (r *repositoryUser) EntityResult(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user models.ModelUser
	user.ID = input.ID

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user, "id=?", input.ID)

	if checkUser.RowsAffected < 1 {
		return &user, checkUser.Error
	}

	return &user, nil
}

func (r *repositoryUser) EntityUpdate(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user models.ModelUser
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Password = input.Password
	user.Role = input.Role

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user, "id=?", input.ID)

	if checkUser.RowsAffected < 1 {
		return &user, checkUser.Error
	}

	updateUser := db.Debug().Save(&user).Commit()

	if updateUser.RowsAffected < 1 {
		return &user, updateUser.Error
	}

	return &user, nil
}

func (r *repositoryUser) EntityDelete(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user models.ModelUser

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user, "id=?", input.ID)

	if checkUser.RowsAffected < 1 {
		return &user, checkUser.Error
	}

	deleteUser := db.Debug().Delete(&user).Commit()

	if deleteUser.RowsAffected < 1 {
		return &user, deleteUser.Error
	}

	return &user, nil
}
