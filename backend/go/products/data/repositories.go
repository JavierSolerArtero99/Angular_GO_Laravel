package data

import (
	"fmt"

	// "time"

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
func FindSingleProduct(name string) (models.Products, error) {
	db := common.GetDB()
	var p models.Products
	var u models.User
	var c []models.Comment

	fmt.Println("---FindSingleProduct---")
	fmt.Println(name)

	// Query
	db.Where("name = ?", name).First(&p)

	if p.Name != "" {
		// Insert all the relationship data
		db.Find(&p).Related(&u, "user")
		db.Find(&p).Related(&c, "comments")

		p.Comments = c
		p.UserModel = u

		// pupulate the comment author
		for i, commentIteration := range p.Comments {
			var commentAuthor models.User
			db.Find(&commentIteration).Related(&commentAuthor, "user")
			p.Comments[i].User = commentAuthor
		}
	}

	return p, nil
}

// Save a product
func SaveComment(commentToSave models.Comment) (models.Comment, error) {
	db := common.GetDB()
	// commentToSave.Date =  time.Now().Unix() + ""
	err := db.Create(&commentToSave).Error

	return commentToSave, err
}

// Save a product
func DeleteComment(commentId int64) (error) {
	db := common.GetDB()
	err := db.Delete(&models.Comment{}, commentId).Error
	return err
}
