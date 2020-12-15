package controllers

import (
	"encoding/json"
	"net/http"

	"products/data"

	"products/common"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
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