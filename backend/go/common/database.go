package common

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("mysql", DbURL(BuildDBConfig()))
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	DB = db
	return DB
}

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "database",
		Port:     3306,
		User:     "root",
		Password: "root",
		DBName:   "thinkster_laravel_ng",
	}
	// fmt.Println(&dbConfig)
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func GetDB() *gorm.DB {
	return DB
}
