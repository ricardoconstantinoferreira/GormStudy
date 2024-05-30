package products

import (
	"gorm/model/products"
	"gorm/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	err := products.Delete(id)

	if err != nil {
		panic(err.Error())
	}

	var message structs.Message
	message.Data = "Product Deleted Successfully"
	c.IndentedJSON(http.StatusOK, message)
}
