package data

import (

	"products/common"
	"products/models"
)

func FindProducts() ([]models.Products, error) {
	db := common.GetDB()
	var model []models.Products
	err := db.Find(&model).Error
	return model, err
	// db := common.GetDB()
	// var model []models.Products
	// tx := db.Begin()
	// tx.Find(&model)
	// tx.Model(&model).Related(&model.Author, "Author")
	// tx.Model(&model.Author).Related(&model.Author)
	// err := tx.Commit().Error
	// return model, err
}