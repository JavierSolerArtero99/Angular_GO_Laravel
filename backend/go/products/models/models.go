package models

import (
)

type (
	Products struct {
		ID		  uint			`gorm:"primary_key" json:"-"`
		Name      string        
		Likes	  int
		User      uint			`json:"-"`
		UserModel User			`gorm:"foreignKey:ID" json:"User"`
		Comments  []Comment		`gorm:"foreignKey:ProductID;references:ID"`
	}
)

type (
	User struct {
		ID           uint    
		Username     string 
	}
)


type (
	Comment struct {
		ID          uint    `gorm:"primary_key"`
		UserID     	uint 	
		ProductID  	uint 	
		Message	   	string
		Date		string
		Likes		int
	}
)