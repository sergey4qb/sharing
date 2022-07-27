package models

type UserRequest struct {
	OpCode int `json:"op_code"`
	Data []User `json:"data"`
}
