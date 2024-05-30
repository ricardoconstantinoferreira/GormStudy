package companies

import (
	"gorm/model/companies"
	"gorm/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	id := c.Param("id")
	companiesModel := structs.Company{}

	if err := c.BindJSON(&companiesModel); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		companies.Update(companiesModel, id)
		c.IndentedJSON(http.StatusOK, companiesModel)
	}
}
