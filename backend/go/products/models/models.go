package models

import (
)

type (
	Products struct {
		// Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name      string        `json:"name"`
		User      uint			`gorm:"foreignKey:ID"`
		UserModel User			
	}
)

type (
	User struct {
		ID           uint    
		Username     string 
	}
)