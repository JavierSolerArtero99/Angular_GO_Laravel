package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"
	// "io/ioutil"
	// "github.com/gin-gonic/gin"


	"products/data"
	"products/models"

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

	// name length control
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

	// find the product
	productModel, err := data.FindSingleProduct(name)

	// product not found
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

func PostComment(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		// errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var c models.Comment

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&c)

	if err != nil {
		fmt.Println("There is errors")
		fmt.Println(err)
		return
	}
		
	savedComment, saveErr := data.SaveComment(c)

	if saveErr != nil {
		fmt.Println("There is errors saving the comment")
		fmt.Println(saveErr)
		fmt.Println(savedComment)
		return
	}

	fmt.Println("---saved Comment---")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("There is errors saving the comment")

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		// errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var c models.Comment

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&c)

	if err != nil {
		fmt.Println("There is errors")
		fmt.Println(err)
		return
	}
		
	deletedComment, deleteErr := data.DeleteComment(c)

	if deleteErr != nil {
		fmt.Println("There is errors saving the comment")
		fmt.Println(deleteErr)
		fmt.Println(deletedComment)
		return
	}

	fmt.Println("---deleted Comment---")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}