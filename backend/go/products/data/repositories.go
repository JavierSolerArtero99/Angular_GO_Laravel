package data

import (
	"fmt"

	"products/common"
	"products/models"
)

func FindProducts() ([]models.Products, error) {
	// db := common.GetDB()
	// var model []models.Products
	// err := db.Find(&model).Error
	// return model, err




	db := common.GetDB()
	var model []models.Products
	tx := db.Begin()
	tx.Find(&model)
	for i, _ := range model {
		tx.Model(model[i]).Related(model[i], "UserModel")
		fmt.Println(model[i])
		// tx.Model(model[i].UserID).Related(models.User,"User")
		// tx.Model().Related(&model.Author, "Author")
	}
	fmt.Println(model)
	err := tx.Commit().Error
	return model, err
}