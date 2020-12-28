package data

import (
	"fmt"

	"products/common"
	"products/models"
)

func FindProducts() ([]models.Products, error) {
	db := common.GetDB()
	var p []models.Products
	var u models.User

	db.Find(&p)
	
	for i, _ := range p {
		u = models.User{ID: p[i].User}
		db.First(&u).Related(&u, "user")
		p[i].UserModel = u
	}
	
	return p, nil
}

func FindSingleProduct() (models.Products, error) {
	db := common.GetDB()
	var p models.Products
	var u models.User
	db.First(&p).Related(&u, "user")
	p.UserModel = u

	return p, nil
}