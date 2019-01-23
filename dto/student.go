package dto

type Student struct {
	Name  string `json:"name"`
	Age   int32  `json:"age"`
	Email string `json:"email"`
	A     Bn     `json:"a"`
}

type Bn struct {
	B int32 `json:"b"`
}
