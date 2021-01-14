package routers

import (
	"github.com/gorilla/mux"
)

// Call all the routers to inicialice them
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router = setProductRouters(router)
	return router
}