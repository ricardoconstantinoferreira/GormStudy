package structs

type Customer struct {
	ID         uint `gorm:"primaryKey;default:auto_random()"`
	Name       string
	Profession string
	Birth      string
	Cpf        string
}
