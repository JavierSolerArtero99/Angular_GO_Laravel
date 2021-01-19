package controllers

import (
	"strconv"
	"encoding/json"
	"fmt"
	"net/http"  

	"github.com/gorilla/mux"

	"products/data"
	"products/models"

	"products/common"
	"github.com/go-redis/redis"
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
	vars := mux.Vars(r)
    slug := vars["id"] // the book title slug
	parsedSlug, err := strconv.ParseInt(slug, 10, 64)
	
	if err != nil {
		msg := "Must introduce an id comment'"
		errorMessage, parsingError := json.Marshal(ProductError{Data: msg})
		if parsingError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(errorMessage)
		return
	}

	deleteErr := data.DeleteComment(parsedSlug)

	if deleteErr != nil {
		msg := "Cannot delete the comment"
		errorMessage, parsingError := json.Marshal(ProductError{Data: msg})
		if parsingError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(errorMessage)
		return
	}

	fmt.Println("---deleted Comment---")
	msg := "Comment deleted"
	successfullDelete, parsingError := json.Marshal(SuccessMessage{Data: msg})
	if parsingError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(successfullDelete)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func BuyProduct(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		// errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	var buy models.Buy

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&buy)

	if err != nil {
		fmt.Println("There is errors")
		fmt.Println(err)
		return
	}

	fmt.Println("BUY")
	fmt.Println(buy)

	result, err := client.Get("buys").Result()
	if err != nil {
		fmt.Println(err)
	}

	if len(result) <= 0 {
		var arrayOfBuys [1]models.Buy
		arrayOfBuys[0] = buy

		json, err := json.Marshal(arrayOfBuys)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("JSON")
		fmt.Println(json)

		client.Set("buys", json, 0)
	} 

	bytes := []byte(result)
 	var buys []models.Buy
	json.Unmarshal(bytes, &buys)
	
	msg := "Buyed"
	successfullDelete, parsingError := json.Marshal(SuccessMessage{Data: msg})
	if parsingError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(successfullDelete)
}
