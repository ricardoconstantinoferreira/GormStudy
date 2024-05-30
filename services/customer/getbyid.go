package customer

import (
	"gorm/model/customer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetById(c *gin.Context) {
	id := c.Param("id")
	customerModel := customer.GetById(id)

	if customerModel.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, customerModel)
	}
}

func GetAll(c *gin.Context) {
	customerModel := customer.GetAll()
	c.IndentedJSON(http.StatusOK, customerModel)
}
