package models

type Request struct {
	OpCode int `json:"op_code"`
	Data []User `json:"data"`
}
