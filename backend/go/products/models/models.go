package models

type (
	Products struct {
		ID          uint `gorm:"primary_key" json:"Id"`
		Slug        string
		Name        string
		Price       float32
		Image       string
		Description string
		Likes       int
		User        uint      `json:"-"`
		UserModel   User      `gorm:"foreignKey:ID" json:"User"`
		Comments    []Comment `gorm:"foreignKey:ProductID;references:ID"`
	}
)

type (
	User struct {
		ID       uint
		Username string
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
