package data

import (
	// "fmt"

	"products/common"
	"products/models"
)

// Finds all products
func FindProducts() ([]models.Products, error) {
	db := common.GetDB()
	var p []models.Products
	
	// Finds all products
	db.Find(&p)
	
	// Loop to insert all the relationship data 
	for i, _ := range p {
		var u models.User
		var c []models.Comment

		// Author
		db.Find(&p[i]).Related(&u, "user")
		p[i].UserModel = u
		// Comments
		db.Find(&p[i]).Related(&c, "comments")
		p[i].Comments = c
	}
	
	return p, nil
}

// Finds a single product
func FindSingleProduct() (models.Products, error) {
	db := common.GetDB()
	var p models.Products
	var u models.User
	var c []models.Comment

	// Query
	db.Where("Name = ?", "Bebida").First(&p)
	// Insert all the relationship data 
	db.Find(&p).Related(&u, "user")
	db.Find(&p).Related(&c, "comments")

	p.Comments = c
	p.UserModel = u

	return p, nil
}