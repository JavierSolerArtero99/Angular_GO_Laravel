package models

import (
	"github.com/jinzhu/gorm"
)

type (
	Products struct {
		gorm.Model
		// Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name      string        `json:"name"`
		User	  int			
		// User      User			
	}
)

type (
	User struct {
		gorm.Model
		ID           int    
		Username     string 
	}
)