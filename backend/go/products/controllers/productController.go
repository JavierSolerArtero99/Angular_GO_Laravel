package controllers

import (
	// "fmt"
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
	// Query params
    name := r.URL.Query()["name"][0] 

	if len(name) <= 0 {
		msg := "The name of the product is empty"
		errorMessage, parsingError := json.Marshal(ProductError{Data: msg})
		if parsingError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(errorMessage)
		return
	}


	productModel, err := data.FindSingleProduct(name)

	if len(productModel.Name) <= 0 {
		msg := "Cannot find product: '" + name + "'"
		errorMessage, parsingError := json.Marshal(ProductError{Data: msg})
		if parsingError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(errorMessage)
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

