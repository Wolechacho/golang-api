package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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

func (o *Order) saveOrder() {
	sess, err := client.StartSession()
	if err != nil {
		log.Fatal("Could not start Session", err)
	}
	err = sess.StartTransaction()
	if err != nil {
		log.Fatal("Could not start Transaction", err)
	}
	err = mongo.WithSession(context.TODO(), sess, func(sc mongo.SessionContext) error {
		collection := client.Database("northwind").Collection("employees")
		employee := Employee{
			ContactName: "Wole Adenigbagbe",
			City:        "Lekki",
			Address:     "Lekki Lagos",
		}
		var empresult *mongo.InsertOneResult
		empresult, err = collection.InsertOne(sc, employee)
		if err != nil {
			sess.AbortTransaction(context.TODO())
			log.Fatal("Could not insert record to employee table", err)
		}

		collection = client.Database("northwind").Collection("customers")
		customer := Customer{
			ContactName: "Ogunyemi Femi",
			City:        "Idumota",
			Address:     "Mainland Lagos",
		}
		var custresult *mongo.InsertOneResult
		custresult, err = collection.InsertOne(sc, customer)
		if err != nil {
			sess.AbortTransaction(context.TODO())
			log.Fatal("Could not insert record to customer table", err)
		}

		var poid primitive.ObjectID
		poid, err = primitive.ObjectIDFromHex("5ddae8d3a4c9080554c1b4e4")
		if err != nil {
			sess.AbortTransaction(context.TODO())
			log.Fatal("Could not convert product id hex of mongodb ObjectId", err)
		}
		items := make([]OrderDetails, 0)
		item := OrderDetails{
			UnitPrice:   40.0,
			Quantity:    4,
			ProductInfo: poid,
			Discount:    2.0,
		}
		items = append(items, item)
		fmt.Printf("Cart to be added %+v : ", items)

		cart := Order{
			OrderDate:    time.Now(),
			ShippedDate:  time.Now(),
			ShipName:     "Adele Voyage",
			ShipAddress:  "Lekki Lagos",
			EmployeeInfo: empresult.InsertedID,
			CustomerInfo: custresult.InsertedID,
			OrderDetails: items,
		}

		collection = client.Database("northwind").Collection("orders")
		var orderresult *mongo.InsertOneResult
		orderresult, err = collection.InsertOne(sc, cart)

		if err != nil {
			sess.AbortTransaction(context.TODO())
			log.Fatal("Could not insert record to order table", err)
		}

		sess.CommitTransaction(context.TODO())
		fmt.Printf("Order Document Inserted with _id : %v", orderresult.InsertedID)
		return nil
	})
	sess.EndSession(context.TODO())
}
