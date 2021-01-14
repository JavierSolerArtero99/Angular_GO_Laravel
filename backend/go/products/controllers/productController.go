package controllers

import (
	"encoding/json"
	"net/http"

	"products/data"

	"products/common"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Get all products from the DB
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Find All products
	productModel, err := data.FindProducts()

	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred, cannot find the product", 500)
		return
	}

	j, err := json.Marshal(ProductsResource{Data: productModel})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Get a product from the DB
func GetSingleProducts(w http.ResponseWriter, r *http.Request) {
	//Finds a single product
	productModel, err := data.FindSingleProduct()

	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred, cannot find the product", 500)
		return
	}

	j, err := json.Marshal(ProductResource{Data: productModel})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler()
}

