package customer

import "gorm/infra"

func Delete(id string) error {
	db := infra.Connect()
	customer := GetById(id)
	tx := db.Delete(&customer)

	return tx.Error
}
