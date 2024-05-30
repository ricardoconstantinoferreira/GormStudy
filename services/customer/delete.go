package customer

import (
	"gorm/model/customer"
	"gorm/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	err := customer.Delete(id)

	if err != nil {
		panic(err.Error())
	}

	var message structs.Message
	message.Data = "Customer Deleted Successfully"
	c.IndentedJSON(http.StatusOK, message)
}
