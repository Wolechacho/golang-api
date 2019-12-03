package models

import (
	"context"
	"fmt"
	"log"

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
func (p *Product) SaveProduct(product *Product) {
	collection := helpers.Client.Database("northwind").Collection("products")

	insertResult, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Product Document: ", insertResult.InsertedID)
}

//GetProducts -- get all products document from the db
func (p *Product) GetProducts() []Product {
	var fetchedProducts []Product
	collection := helpers.Client.Database("northwind").Collection("products")

	cur, err := collection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var product Product
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}
		fetchedProducts = append(fetchedProducts, product)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return fetchedProducts
}

//GetProductByID -- get a product document with the hex id
func (p *Product) GetProductByID(id string) Product {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	var product Product
	collection := helpers.Client.Database("northwind").Collection("products")
	err = collection.FindOne(context.TODO(), bson.M{"_id": oid}, options.FindOne()).Decode(&product)
	if err != nil {
		log.Fatal(err)
	}
	return product
}

//DeleteProductByID -- delete a product document with hex id
func (p *Product) DeleteProductByID(id string) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	collection := helpers.Client.Database("northwind").Collection("products")
	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": oid}, options.Delete())
	if err != nil {
		log.Fatal(err)
	}
	if result.DeletedCount != 1 {
		log.Fatal(err)
	}
	fmt.Printf("Number of product deleted  : %d\n", result.DeletedCount)
}
