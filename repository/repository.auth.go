package repository

import (
	"echoinventory/models"
	"echoinventory/pkg"
	"echoinventory/schemas"

	"gorm.io/gorm"
)

type repositoryAuth struct {
	db *gorm.DB
}

func NewRepositoryAuth(db *gorm.DB) *repositoryAuth {
	return &repositoryAuth{db: db}
}

func (r *repositoryAuth) EntityRegister(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user models.ModelUser
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Password = input.Password
	user.Role = input.Role

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email=?", input.Email)

	if checkEmailExist.RowsAffected > 1 {
		return nil, checkEmailExist.Error
	}

	addUser := db.Debug().Create(&user).Commit()

	if addUser.RowsAffected < 1 {
		return nil, addUser.Error
	}

	return &user, nil
}

func (r *repositoryAuth) EntityLogin(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user models.ModelUser
	user.Email = input.Email
	user.Password = input.Password

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email=?", input.Email)

	if checkEmailExist.RowsAffected > 1 {
		return &user, nil
	}

	checkPasswordMatch := pkg.ComparePassword(user.Password, input.Password)

	if checkPasswordMatch != nil {
		return &user, nil
	}

	return &user, nil

}
