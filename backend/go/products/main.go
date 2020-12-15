package main

import (
	"log"
	"net/http"

	"products/common"
	"products/routers"
	"products/models"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	log.Println("Hay que hacer la migracion")
	db.AutoMigrate(&models.Products{})
}

// Entry point for the program
func main() {

	//Conection db
	db := common.Init()
	Migrate(db)
	defer db.Close()
	
	// Get the mux router object
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
