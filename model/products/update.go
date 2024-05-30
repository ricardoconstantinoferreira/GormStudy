package products

import (
	"gorm/infra"
	"gorm/structs"
)

func Update(productsModel structs.Product, id string) {
	db := infra.Connect()
	product := GetById(id)

	product.Code = productsModel.Code
	product.Price = productsModel.Price

	db.Save(&product)
}
