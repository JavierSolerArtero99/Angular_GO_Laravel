package routers

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	
	"github.com/gorilla/mux"
	"products/controllers"
)

// Inicialice the products routes
func setProductRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/products/", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products/product", controllers.GetSingleProducts).Methods("GET")
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