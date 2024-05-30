package products

import "gorm/infra"

func Delete(id string) error {
	db := infra.Connect()
	product := GetById(id)
	tx := db.Delete(&product)

	return tx.Error
}
