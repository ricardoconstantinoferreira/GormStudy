package companies

import (
	"gorm/infra"
	"gorm/structs"
)

func Update(companiesModel structs.Company, id string) {
	db := infra.Connect()
	company := GetById(id)

	company.Name = companiesModel.Name
	company.Fantasy_name = companiesModel.Fantasy_name
	company.Cnpj = companiesModel.Cnpj

	db.Save(&company)
}
