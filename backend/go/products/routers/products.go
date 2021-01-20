package routers

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"products/controllers"

	"github.com/gorilla/mux"
)

// Inicialice the products routes
func setProductRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products/product", controllers.GetSingleProducts).Methods("GET")
	router.HandleFunc("/products/like/{name}", controllers.LikeProduct).Methods("POST")
	router.HandleFunc("/products/comment", controllers.PostComment).Methods("POST")
	router.HandleFunc("/products/buy", controllers.ProductsBuys).Methods("GET")
	router.HandleFunc("/products/buy", controllers.BuyProduct).Methods("POST")
	router.HandleFunc("/products/comment/{id}", controllers.DeleteComment).Methods("DELETE")
	recordMetrics()
	router.HandleFunc("/metrics", controllers.GetMetrics).Methods("GET")
	return router
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)
