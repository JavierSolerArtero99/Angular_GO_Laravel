package models

import (
)

type (
	Products struct {
		ID		  uint			`gorm:"primary_key"`
		Name      string        
		Likes	  int
		User      uint			
		UserModel User			`gorm:"foreignKey:ID"`
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
		ProductID  	string 	
		Message	   	string
		Date		string
		Likes		int
	}
)