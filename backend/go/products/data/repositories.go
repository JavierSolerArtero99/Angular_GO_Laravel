package data

import (
	"fmt"

	"products/common"
	"products/models"
)

func FindProducts() ([]models.Products, error) {
	db := common.GetDB()
	var p []models.Products
	
	db.Find(&p)
	
	for i, _ := range p {
		var u models.User
		var c []models.Comment

		db.Find(&p[i]).Related(&u, "user")
		p[i].UserModel = u
		db.Find(&p[i]).Related(&c, "comments")
		p[i].Comments = c

		fmt.Println(u)
	}
	
	return p, nil
}

func FindSingleProduct() (models.Products, error) {
	db := common.GetDB()
	var p models.Products
	var u models.User
	var c []models.Comment

	db.Where("Name = ?", "Consumible").First(&p)
	db.Find(&p).Related(&u, "user")
	db.Find(&p).Related(&c, "comments")

	p.Comments = c
	p.UserModel = u

	return p, nil
}