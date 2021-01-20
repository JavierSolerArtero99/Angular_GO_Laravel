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

	// Control del body de la peticion
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		// errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	// Conexion con redis
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	// Decoding body
	var buy models.Buy

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&buy)

	if err != nil {
		fmt.Println("There is errors")
		fmt.Println(err)
		return
	}

	// Obteniendo los resultados de redis
	result, err := client.Get("buys").Result()
	if err != nil {
		fmt.Println(err)
	}

	// Si la estructura solicitada no existe en redis se creará
	if len(result) <= 0 {
		var arrayOfBuys [1]models.Buy
		buy.TimesBuyed = 1	// se pone en 1 la cantidad de veces comprado
		arrayOfBuys[0] = buy

		json, err := json.Marshal(arrayOfBuys)
		if err != nil {
			fmt.Println(err)
		}

		client.Set("buys", json, 0)
	
	} else {
		// En caso contrario se añadira la compra
		bytes := []byte(result)	// Leyendo las compras de redis
		var buys []models.Buy
		var isInside bool = false
		json.Unmarshal(bytes, &buys)
		
		/* Comprobando si la nueva compra esta dentro de las compras ya hechas, 
		si está dentro se actualizará el contador de la compra */
		for i, iterationBuy := range buys {
			if buy.Product == iterationBuy.Product {
				isInside = true
				buys[i].TimesBuyed++
			}
		}

		if isInside {
			// Actualizando los datos de la compra
			newBuysJSON, err := json.Marshal(buys)
			if err != nil {
				fmt.Println(err)
			}
			client.Set("buys", newBuysJSON, 0)

		} else {
			// Creando la nueva compra 
			newLenght := len(buys) + 1
			var newBuys = make([]models.Buy,(newLenght))
			buy.TimesBuyed = 1

			for i := range buys {
				newBuys[i] = buys[i]
			}

			newBuys[len(buys)] = buy
			fmt.Println(newBuys)

			newBuysJSON, err := json.Marshal(newBuys)
			if err != nil {
				fmt.Println(err)
			}
			client.Set("buys", newBuysJSON, 0)
		}

	}
	
	msg := "Buyed"
	successfullBuyed, parsingError := json.Marshal(SuccessMessage{Data: msg})
	if parsingError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(successfullBuyed)
}

func ProductsBuys(w http.ResponseWriter, r *http.Request) {
	// Conexion con redis
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	
	// Obteniendo los resultados de redis
	result, err := client.Get("buys").Result()
	if err != nil {
		fmt.Println(err)
	}

	// Parsing compras de string a json
	bytes := []byte(result)	
	var buys []models.Buy
	json.Unmarshal(bytes, &buys)
	buysJSON, err := json.Marshal(BuyStats{Data: buys})

	// Enviando respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buysJSON)
}

func LikeProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    name := vars["name"] // the book title slug

	// Augmentando en uno el like del producto
	likesAmount := productModel.Likes + 1
	err := data.LikeProduct(name, likesAmount)
	// Error cuando no se ha podido hacer un like
	if err != nil {
		msg := "Cannot like the product: '" + name + "'"
		errorMessage, parsingError := json.Marshal(ProductError{Data: msg})
		if parsingError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(errorMessage)
		return
	}

	// Enviando respuesta
	msg := "Product liked"
	successfullLiked, parsingError := json.Marshal(SuccessMessage{Data: msg})
	if parsingError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(successfullLiked)
}
