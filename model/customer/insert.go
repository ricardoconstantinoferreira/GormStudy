package customer

import (
	"gorm/infra"
	"gorm/structs"
)

func Insert(customerModel structs.Customer) {
	db := infra.Connect()
	db.Create(&customerModel)
}
