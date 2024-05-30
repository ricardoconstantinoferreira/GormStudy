package structs

type Product struct {
	ID    uint `gorm:"primaryKey;default:auto_random()"`
	Code  string
	Price int
}
