package models

type User struct {
	ID      int    `json:"id"`
	Code    string `json:"code"`
	Age     int    `json:"age"`
	Name    string `json:"name"`
	Height  int    `json:"height"`
	Address string `json:"address"`
}
