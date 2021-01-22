package data

import (
	"fmt"
	"strconv"
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
		var l []models.LikeList

		// Author
		db.Find(&p[i]).Related(&u, "user")
		p[i].UserModel = u
		// Comments
		db.Find(&p[i]).Related(&c, "comments")
		p[i].Comments = c
		// Likes
		db.Where("product_id = ?", p[i].ID).Find(&l)
		p[i].LikesUserList = l
	}

	return p, nil
}

// Finds a single product
func FindSingleProduct(name string) (models.Products, error) {
	db := common.GetDB()
	var p models.Products
	var u models.User
	var c []models.Comment
	var l []models.LikeList

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

		// Likes
		db.Where("product_id = ?", p.ID).Find(&l)
		p.LikesUserList = l
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

func LikeProduct(productToLike string, userId string) error {
	db := common.GetDB()
	var p models.Products

	// Query
	db.Where("name = ?", productToLike).First(&p)
	err := db.Model(&p).Update("likes", p.Likes + 1).Error

	i2, err := strconv.ParseInt(userId, 10, 64)
	if err == nil {
		fmt.Println(i2)
	}

	if err == nil {
		likeList := models.LikeList{
			ProductID: p.ID,
			UserID: i2,
		}
		err = db.Create(&likeList).Error
	}


	return err
}

// Save a product
func DeleteComment(commentId int64) (error) {
	db := common.GetDB()
	err := db.Delete(&models.Comment{}, commentId).Error
	return err
}
