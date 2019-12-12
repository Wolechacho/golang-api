package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//OrderDetails contain product information as well as quantity
type OrderDetails struct {
	UnitPrice   float64            `json:"unitPrice"`
	Quantity    int                `json:"quantity"`
	Discount    float64            `json:"discount"`
	ProductInfo primitive.ObjectID `json:"productInfo"`
}
