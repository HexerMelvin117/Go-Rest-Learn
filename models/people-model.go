package models

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
}
