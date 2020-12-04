package products

import (
	"App/common"
	"gopkg.in/gin-gonic/gin.v1"
)

// *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
type ProductModelValidator struct {
	Product struct {
		Name 	 string `form:"name" json:"name" binding:"exists,alphanum,min=4,max=255"`
	} `json:"product"`
	productModel ProductModel `json:"-"`
}

// There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (self *ProductModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.productModel.Name = self.Product.Name

	return nil
}

// You can put the default value of a Validator here
func NewProductModelValidator() ProductModelValidator {
	productModelValidator := ProductModelValidator{}
	//productModelValidator.Product.Email ="w@g.cn"
	return productModelValidator
}

func NewProductModelValidatorFillWith(productModel ProductModel) ProductModelValidator {
	productModelValidator := NewProductModelValidator()
	productModelValidator.Product.Name = productModel.Name

	return productModelValidator
}