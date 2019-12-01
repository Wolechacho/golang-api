package models

import 	"time"

//Order model specifies the cart information
type Order struct {
	OrderDate    time.Time
	ShippedDate  time.Time
	ShipName     string
	ShipAddress  string
	OrderDetails []OrderDetails
	EmployeeInfo interface{}
	CustomerInfo interface{}
}