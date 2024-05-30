package customer

import (
	"gorm/infra"
	"gorm/structs"
)

func GetById(id string) *structs.Customer {
	db := infra.Connect()

	customerModel := structs.Customer{}
	db.Where("id = ?", id).First(&customerModel)
	return &customerModel
}

func GetAll() *[]structs.Customer {
	db := infra.Connect()

	var customerModel []structs.Customer
	db.Find(&customerModel).Scan(&customerModel)
	return &customerModel
}
