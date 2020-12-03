package products

import (
	"fmt"

	"github.com/javiersoler/Angular_GO_Laravel/backend/go/common"
)

/*------PRODUCT MODEL------*/
type ProductModel struct {
	ID           uint    `gorm:"primary_key"`
	Name     	 string  `gorm:"column:name; not null"`
	// Email        string  `gorm:"column:email;unique_index"`
	// Bio          string  `gorm:"column:bio;size:1024"`
	// Image        *string `gorm:"column:image"`
	// PasswordHash string  `gorm:"column:password;not null"`
	// Role 		 bool 	 `gorm:"column:role;default 0 not null"`
}

/** Migrates the database */
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&ProductModel{})
}


/**------DATABASE OPERATIONS------**/

/* Finds all products */
func FindAllProducts() ([]ProductModel, error) {
	db := common.GetDB()
	var models []ProductModel
	err := db.Find(&models).Error

	return models, err
}

/* Finds a product */
func FindOneProduct(condition interface{}) (ProductModel, error) {
	db := common.GetDB()
	var model ProductModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

/* Save a product */
func SaveOne(data interface{}) error {
	fmt.Println(data)
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}