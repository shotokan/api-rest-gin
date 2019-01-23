package models

type Student struct {
	User
	Grade   string
	A       Bn
	Address Address
	BaseModel
}

type Bn struct {
	B int32 `json:"b"`
}
