package products

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type ProductSerializer struct {
	c *gin.Context
	ProductModel
}

type ProductsSerializer struct {
	c *gin.Context
	Products []ProductModel
}

type ProductResponse struct {
	Name 		string `json:"name"`
	// Email    string  `json:"email"`
	// Bio      string  `json:"bio"`
	// Image    *string `json:"image"`
	// Token    string  `json:"token"`
	// Role	 bool	 `json:"role"`
}

func (self *ProductSerializer) Response() ProductResponse {
	// myProductModel := self.c.MustGet("my_product_model").(ProductModel)
	product := ProductResponse{
		Name:        self.Name,
		// Email:    myUserModel.Email,
		// Bio:      myUserModel.Bio,
		// Image:    myUserModel.Image,
		// Token:    common.GenToken(myUserModel.ID),
		// Role:     myUserModel.Role,
	}
	return product
}


func (s *ProductsSerializer) Response() []ProductResponse {
	response := []ProductResponse{}
	for _, tag := range s.Products {
		serializer := ProductSerializer{s.c, tag}
		response = append(response, serializer.Response())
	}
	return response
}
