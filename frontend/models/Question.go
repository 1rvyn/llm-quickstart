package models

type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Answer   *string  `json:"answer"`
	User     Accounts `json:"user"`
}
