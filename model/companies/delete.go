package companies

import "gorm/infra"

func Delete(id string) error {
	db := infra.Connect()
	companiesModel := GetById(id)
	tx := db.Delete(&companiesModel)

	return tx.Error
}
