package products

import (
	"gorm/infra"
	"gorm/structs"
)

func GetById(id string) *structs.Product {
	db := infra.Connect()

	productsModel := structs.Product{}
	db.Where("id = ?", id).First(&productsModel)
	return &productsModel
}

func GetAll() *[]structs.Product {
	db := infra.Connect()

	productsModel := []structs.Product{}
	db.Find(&productsModel).Scan(&productsModel)
	return &productsModel
}
