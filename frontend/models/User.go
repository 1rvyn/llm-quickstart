package models

type Accounts struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password []byte `gorm:"not null"`
	UserRole int    `gorm:"default:0"`
}
