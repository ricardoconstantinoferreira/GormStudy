package products

import (
	"gorm/model/products"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetById(c *gin.Context) {
	id := c.Param("id")
	productsModel := products.GetById(id)

	if productsModel.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, productsModel)
	}
}

func GetAll(c *gin.Context) {
	productsModel := products.GetAll()
	c.IndentedJSON(http.StatusOK, productsModel)
}
