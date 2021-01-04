package main

import (
	"log"
	"net/http"

	"products/common"
	"products/routers"
	"products/models"
	"github.com/jinzhu/gorm"
)

/** Migrates the database schema */
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Products{})
	db.AutoMigrate(&models.Comment{})
}

// Entry point for the program
func main() {

	//Conection db
	db := common.Init()
	Migrate(db)
	defer db.Close()
	
	// Get the mux router object
	router := routers.InitRoutes()

	// Init server
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
