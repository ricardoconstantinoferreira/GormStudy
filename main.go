package main

import (
	"gorm/infra"
	"gorm/services/companies"
	"gorm/services/customer"
	"gorm/services/products"
	"gorm/structs"

	"github.com/gin-gonic/gin"
)

func main() {

	createTables()

	router := gin.Default()

	//Customers endpoints
	router.POST("/customer", customer.Insert)
	router.GET("/customer/:id", customer.GetById)
	router.GET("/customer", customer.GetAll)
	router.PUT("/customer/:id", customer.Update)
	router.DELETE("/customer/:id", customer.Delete)

	//Companies endpoints
	router.POST("/companies", companies.Insert)
	router.GET("/companies/:id", companies.GetById)
	router.GET("/companies", companies.GetAll)
	router.PUT("/companies/:id", companies.Update)
	router.DELETE("/companies/:id", companies.Delete)

	//Products endpoints
	router.POST("/products", products.Insert)
	router.GET("/products/:id", products.GetById)
	router.GET("/products", products.GetAll)
	router.PUT("/products/:id", products.Update)
	router.DELETE("/products/:id", products.Delete)

	router.Run("localhost:8080")

}

func createTables() {
	db := infra.Connect()
	db.AutoMigrate(&structs.Product{})
	db.AutoMigrate(&structs.Customer{})
	db.AutoMigrate(&structs.Company{})
}
