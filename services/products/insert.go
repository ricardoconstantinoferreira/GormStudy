package products

import (
	"gorm/model/products"
	"gorm/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context) {
	productsModel := structs.Product{}

	if err := c.BindJSON(&productsModel); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		products.Insert(productsModel)
		c.IndentedJSON(http.StatusCreated, productsModel)
	}
}
