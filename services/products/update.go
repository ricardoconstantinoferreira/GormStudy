package products

import (
	"gorm/model/products"
	"gorm/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	id := c.Param("id")
	productsModel := structs.Product{}

	if err := c.BindJSON(&productsModel); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		products.Update(productsModel, id)
		c.IndentedJSON(http.StatusCreated, productsModel)
	}
}
