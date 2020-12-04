package products

import (
	"fmt"

	"App/common"
)

type ProductModel struct {
	ID           uint    `gorm:"primary_key"`
	Name     	 string  `gorm:"column:name; not null"`
	// Email        string  `gorm:"column:email;unique_index"`
	// Bio          string  `gorm:"column:bio;size:1024"`
	// Image        *string `gorm:"column:image"`
	// PasswordHash string  `gorm:"column:password;not null"`
	// Role 		 bool 	 `gorm:"column:role;default 0 not null"`
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&ProductModel{})
}

func FindAllProducts() ([]ProductModel, error) {
	db := common.GetDB()
	var models []ProductModel
	err := db.Find(&models).Error

	return models, err
}

func FindOneProduct(condition interface{}) (ProductModel, error) {
	db := common.GetDB()
	var model ProductModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

func SaveOne(data interface{}) error {
	fmt.Println(data)
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func (model *ProductModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Update(data).Error
	return err
}