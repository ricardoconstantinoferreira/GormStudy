package customer

import (
	"gorm/model/customer"
	"gorm/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	id := c.Param("id")
	customerModel := structs.Customer{}

	if err := c.BindJSON(&customerModel); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		customer.Update(customerModel, id)
		c.IndentedJSON(http.StatusCreated, customerModel)
	}

}
