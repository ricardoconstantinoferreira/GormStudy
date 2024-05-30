package companies

import (
	"gorm/model/companies"
	"gorm/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	err := companies.Delete(id)

	if err != nil {
		panic(err.Error())
	}

	var message structs.Message
	message.Data = "Company Deleted Successfully"
	c.IndentedJSON(http.StatusOK, message)
}
