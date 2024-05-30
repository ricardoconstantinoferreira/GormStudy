package companies

import (
	"gorm/model/companies"
	"gorm/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context) {
	companiesModel := structs.Company{}

	if err := c.BindJSON(&companiesModel); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		companies.Insert(companiesModel)
		c.IndentedJSON(http.StatusOK, companiesModel)
	}
}
