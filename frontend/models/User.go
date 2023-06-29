package models

import "time"

type Accounts struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `json:"Username" gorm:"unique;unique_index"`
	Password  []byte
	UserRole  int       `gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
}
