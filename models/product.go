package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Product model contains a product information
type Product struct {
	ProductName  string
	UnitPrice    float64
	UnitInStock  int
	Discontinued bool
	CategoryInfo primitive.ObjectID
}
