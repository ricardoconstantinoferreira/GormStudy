package companies

import (
	"gorm/infra"
	"gorm/structs"
)

func Insert(companiesModel structs.Company) {
	db := infra.Connect()
	db.Create(&companiesModel)
}
