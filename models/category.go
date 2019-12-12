package models

import (
	"context"
	"first-api-golang/helpers"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var err error

//Category groups the product
type Category struct {
	CategoryName string `json:"categoryName"`
	Description  string `json:"description"`
}

//SaveCategory -- save categories of product into the DB
func (c *Category) SaveCategory(category *Category) error {
	collection := helpers.Client.Database("northwind").Collection("categories")

	newcategory := Category{CategoryName: category.CategoryName, Description: category.Description}
	insertResult, err := collection.InsertOne(context.TODO(), newcategory)

	if err != nil {
		return fmt.Errorf("Cagetory not saved %+v", err)
	}
	fmt.Println("Inserted a Single Category Document : ", insertResult.InsertedID)
	return nil
}
