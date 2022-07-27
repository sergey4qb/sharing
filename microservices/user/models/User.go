package models


type User struct {
	ID      uint32    `json:"id" `
	Name      string    `json:"name" `
	Address   string    `json:"address" `
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"update_at"`
}



