package models

import (
)

type (
	Products struct {
		// Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name      string        `json:"name"`
		Author	  int			`json: user`
	}
)

type (
	Author struct {
		ID           uint    `gorm:"primary_key"`
		Username     string  `gorm:"column:username"`
		Email        string  `gorm:"column:email;unique_index"`
		Bio          string  `gorm:"column:bio;size:1024"`
		Image        *string `gorm:"column:image"`
		PasswordHash string  `gorm:"column:password;not null"`
		Role         bool    `gorm:"column:role;default 0 not null"`
		Karma        uint    `gorm:"column:karma;default 0 not null"`
		Tempkey      string  `gorm:"column:tempkey`
	}
)