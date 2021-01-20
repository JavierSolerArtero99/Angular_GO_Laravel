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

	BuyStats struct {
		Data []models.Buy `json:"buys"`
	}

	ProductError struct {
		Data string `json:"error"`
	}

	SuccessMessage struct {
		Data string `json:"success"`
	}
)
