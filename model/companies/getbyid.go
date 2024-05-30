package companies

import (
	"gorm/infra"
	"gorm/structs"
)

func GetById(id string) *structs.Company {
	db := infra.Connect()

	companiesModel := structs.Company{}
	db.Where("id = ?", id).First(&companiesModel)
	return &companiesModel
}

func GetAll() *[]structs.Company {
	db := infra.Connect()

	companiesModel := []structs.Company{}
	db.Find(&companiesModel).Scan(&companiesModel)
	return &companiesModel
}
