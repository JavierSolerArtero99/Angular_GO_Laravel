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
}