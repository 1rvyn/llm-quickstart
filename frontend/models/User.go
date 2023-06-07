package models

type Users struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password []byte `gorm:"not null"`
}
