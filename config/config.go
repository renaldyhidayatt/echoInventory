package config

import (
	"echoinventory/models"
	"echoinventory/pkg"
	"fmt"

	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", pkg.GodotEnv("DB_HOST"), pkg.GodotEnv("DB_PORT"), pkg.GodotEnv("DB_USER"), pkg.GodotEnv("DB_NAME"), pkg.GodotEnv("DB_PASSWORD"))

	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
		defer logger.Info("Database connection failed")
		logger.Fatal(err)
		return nil
	}

	err = db.AutoMigrate(&models.ModelCategory{}, &models.ModelUser{}, &models.ModelSupplier{}, &models.ModelCustomer{}, &models.ModelProduct{}, &models.ModelProductKeluar{}, &models.ModelProductMasuk{})

	if err != nil {
		defer logger.Info("Database connection failed")
		logger.Fatal(err)

		return nil
	}

	return db

}
