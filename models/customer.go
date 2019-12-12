package models

//Customer model specifies information about the buyer
type Customer struct {
	ContactName string `json:"contactname"`
	Address     string `json:"address"`
	City        string `json:"city"`
}
