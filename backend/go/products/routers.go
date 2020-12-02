package products

import (
	"fmt"

	"net/http"

	"github.com/javiersoler/Angular_GO_Laravel/backend/go/common"
	"gopkg.in/gin-gonic/gin.v1"
)



/* --------ROUTER: Get products register-------- */

func ProductsRegister(router *gin.RouterGroup) {
	router.GET("/", ProductRetrieve)
	router.POST("/", ProductCreate)
}

/* CONTROLERS */

func ProductRetrieve(p *gin.Context) {
	// fmt.Println("Entra en el controlador")
	// productModel, err := FindProducts(&ProductModel)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid slug")))
	// 	return
	// }

	// serializer := ProductSerializer{p, ProductModel}
	// p.JSON(http.StatusOK, gin.H{"ESSE": serializer.Response()})
}

func ProductCreate(c *gin.Context) {
	productModelValidator := NewProductModelValidator()
	if err := productModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}


	if err := SaveOne(&productModelValidator.productModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := ProductSerializer{c, productModelValidator.productModel}
	fmt.Println("HSTA AKII")
	c.JSON(http.StatusCreated, gin.H{"product": serializer.Response()})
}