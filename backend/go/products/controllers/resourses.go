package controllers

import (
	"products/models"
)

type (
	// For Get - /products
	ProductsResource struct {
		Data []models.Products `json:"products"`
	}
	// For Post/Put - /products
	ProductResource struct {
		Data models.Products `json:"product"`
	}

	ProductError struct {
		Data string `json:"error"`
	}
)
