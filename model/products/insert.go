package products

import (
	"gorm/infra"
	"gorm/structs"
)

func Insert(productsModel structs.Product) {
	db := infra.Connect()
	db.Create(&productsModel)
}
