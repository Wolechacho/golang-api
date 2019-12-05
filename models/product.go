package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"first-api-golang/helpers"
)

//Product model contains a product information
type Product struct {
	ProductName  string             `json:"productName"`
	UnitPrice    float64            `json:"unitPrice"`
	UnitInStock  int                `json:"unitInStock"`
	Discontinued bool               `json:"discontinued"`
	CategoryInfo primitive.ObjectID `json:"categoryInfo"`
}

//SaveProduct -- insert a new product to the db
func (p *Product) SaveProduct(product *Product) error {
	collection := helpers.Client.Database("northwind").Collection("products")

	insertResult, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		return fmt.Errorf("Product not saved %v", err)
	}
	fmt.Println("Inserted a Single Product Document: ", insertResult.InsertedID)
	return nil
}

//GetProducts -- get all products document from the db
func (p *Product) GetProducts() ([]Product, error) {
	var fetchedProducts []Product
	collection := helpers.Client.Database("northwind").Collection("products")

	cur, err := collection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		return nil, fmt.Errorf("Error retrieving products %v", err)
	}
	for cur.Next(context.TODO()) {
		var product Product
		_ = cur.Decode(&product)
		fetchedProducts = append(fetchedProducts, product)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Error retrieving products %v", err)
	}
	cur.Close(context.TODO())
	return fetchedProducts, nil
}

//GetProductByID -- get a product document with the hex id
func (p *Product) GetProductByID(id string) (interface{}, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Could not convert hex id : %v", err)
	}
	var product Product
	collection := helpers.Client.Database("northwind").Collection("products")
	err = collection.FindOne(context.TODO(), bson.M{"_id": oid}, options.FindOne()).Decode(&product)
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve the  product : %v", err)
	}
	return product, nil
}

//DeleteProductByID -- delete a product document with hex id
func (p *Product) DeleteProductByID(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("could not convert hex _id : %v", err)
	}
	collection := helpers.Client.Database("northwind").Collection("products")
	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": oid}, options.Delete())
	if result.DeletedCount != 1 {
		return fmt.Errorf("could not delete the product : %v", err)
	}
	return nil
}
