package products

import (
	"fmt"

	"errors"

	"net/http"

	"github.com/javiersoler/Angular_GO_Laravel/backend/go/common"
	"gopkg.in/gin-gonic/gin.v1"
)


/* --------ROUTER: Get products register-------- */

func ProductsRegister(router *gin.RouterGroup) {
	router.GET("/", ProductList)
	router.POST("/", ProductCreate)
}

/* CONTROLERS */

func ProductList(p *gin.Context) {
	productModels, err := FindAllProducts()
	if err != nil {
		p.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid param")))
		return
	}
	serializer := ProductsSerializer{p, productModels}
	p.JSON(http.StatusOK, gin.H{"products": serializer.Response()})
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