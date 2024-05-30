package customer

import (
	"gorm/infra"
	"gorm/structs"
)

func Update(customerModel structs.Customer, id string) {
	db := infra.Connect()
	customer := GetById(id)

	customer.Name = customerModel.Name
	customer.Profession = customerModel.Profession
	customer.Cpf = customerModel.Cpf
	customer.Birth = customerModel.Birth

	db.Save(&customer)
}
