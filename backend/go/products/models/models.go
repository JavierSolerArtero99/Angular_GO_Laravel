package models

type (
	Products struct {
		ID        uint `gorm:"primary_key" json:"Id"`
		Name      string
		Image      string
		Likes     int
		Price	  uint	 	`json:"Price"`
		User      uint      `json:"-"`
		UserModel User      `gorm:"foreignKey:ID" json:"User"`
		Comments  []Comment `gorm:"foreignKey:ProductID;references:ID"`
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
		Product		string	
		Price 		int			
		TimesBuyed	int
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
