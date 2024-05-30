package structs

type Company struct {
	ID           uint `gorm:"primaryKey;default:auto_random()"`
	Name         string
	Fantasy_name string
	Cnpj         string
}
