package companies

import (
	"gorm/model/companies"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetById(c *gin.Context) {
	id := c.Param("id")
	companiesModel := companies.GetById(id)

	if companiesModel.ID == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.IndentedJSON(http.StatusOK, companiesModel)
	}
}

func GetAll(c *gin.Context) {
	companiesModel := companies.GetAll()
	c.IndentedJSON(http.StatusOK, companiesModel)
}
