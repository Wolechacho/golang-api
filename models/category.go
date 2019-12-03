package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var err error

//Category groups the product
type Category struct {
	CategoryName string
	Description  string
}

func (c *Category) saveCategory() {
	collection := client.Database("northwind").Collection("categories")

	newcategory := Category{CategoryName: "Kitchen Utensils", Description: "Used for house chores"}

	insertResult, err := collection.InsertOne(context.TODO(), newcategory)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Category Document : ", insertResult.InsertedID)
}
