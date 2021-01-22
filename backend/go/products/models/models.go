package models

import "time"

type (
	Products struct {
		ID          uint `gorm:"primary_key" json:"Id"`
		Slug        string
		Name        string
		Price       float32
		Image       string
		Description string
		Likes       int
		LikesUserList	[]LikeList	 `gorm:"many2many:like_lists;" json:"LikesList"`
		CreatedAt   time.Time `sql:",null"`
		UpdatedAt   time.Time `sql:",null"`
		UserID      uint      `json:"-"`
		UserModel   User      `gorm:"foreignKey:ID" json:"User"`
		Comments    []Comment `gorm:"foreignKey:ProductID;references:ID"`
	}
)

type (
	LikeList struct {
		ProductID uint
		UserID    int64
	}
)

type (
	User struct {
		ID       uint
		Username string
	}
)

type (
	Buy struct {
		Product    string
		Price      int
		TimesBuyed int
	}
)

type (
	Comment struct {
		ID        uint   `gorm:"primary_key" json:"Id`
		UserID    uint   `gorm:"foreignKey:ID" json:"UserID"`
		User      User   `json:"Author"`
		ProductID uint   `json:"ProductID"`
		Message   string `json:"Message"`
		Date      string `json:"Date"`
		Likes     int    `json:"Likes"`
	}
)
