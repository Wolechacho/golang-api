package models

//Employee model specifies who sold the product
type Employee struct {
	ContactName string `json:"contactname"`
	Address     string `json:"address"`
	City        string `json:"city"`
}
