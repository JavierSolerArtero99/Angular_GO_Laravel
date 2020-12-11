package controllers

import (
	"fmt"
	
	"encoding/json"
	"net/http"

	// "github.com/gorilla/mux"
	// "products/common"
	// // "products/data"
	// "gopkg.in/mgo.v2"

	// "fmt"
	// "os"
	// "github.com/garyburd/redigo/redis"
)

type Message struct {
    Name string
    Body string
}

// Handler for HTTP Get - "/movies"
// Returns all Movie documents
func GetProducts(w http.ResponseWriter, r *http.Request) {
	m := Message{"Hello", "World"}
	j, err := json.Marshal(m)
	if err != nil {
		fmt.Println("ERROOOOOOOOOOOOOOOOOOOOOOOOOOOR")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}