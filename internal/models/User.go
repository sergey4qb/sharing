package models

type User struct {
	AdminID int `json:"admin_id"`
	ID int `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Balance  int    `json:"balance"`
}
