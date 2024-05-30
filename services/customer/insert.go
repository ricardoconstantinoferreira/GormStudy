package customer

import (
	"gorm/model/customer"
	"gorm/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context) {
	customerModel := structs.Customer{}

	if err := c.BindJSON(&customerModel); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		customer.Insert(customerModel)
		c.IndentedJSON(http.StatusCreated, customerModel)
	}
}
