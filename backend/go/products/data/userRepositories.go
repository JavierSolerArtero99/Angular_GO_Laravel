package data

import (

	"products/common"
	"products/models"
)

func FindAuthor() ([]models.Products, error) {
	db := common.GetDB()
	var model []models.Products
	err := db.Find(&model).Error
	return model, err
}