package models

type Question struct {
	ID       int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Question string  `json:"question" gorm:"not null"`
	Answer   *string `json:"answer"`
	UserID   uint    `json:"user"`
}
