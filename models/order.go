package models

import (
	"context"
	"first-api-golang/helpers"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//Order model specifies the cart information
type Order struct {
	OrderDate        time.Time      `json:"orderDate"`
	ShippedDate      time.Time      `json:"shippedDate"`
	ShipName         string         `json:"shipName"`
	ShipAddress      string         `json:"shipAddress"`
	OrderDetailsList []OrderDetails `json:"orderDetails"`
	Employee         Employee       `json:"employee"`
	Customer         Customer       `json:"customer"`
}

//SaveOrder -- Insert customer order
func (o *Order) SaveOrder(order *Order) error {
	sess, err := helpers.Client.StartSession()
	if err != nil {
		return fmt.Errorf("Could not start Session : %+v", err)
	}
	err = sess.StartTransaction()
	if err != nil {
		return fmt.Errorf("Could not start Transaction : %+v", err)
	}
	err = mongo.WithSession(context.TODO(), sess, func(sc mongo.SessionContext) error {
		collection := helpers.Client.Database("northwind").Collection("employees")
		employee := Employee{
			ContactName: order.Employee.ContactName,
			City:        order.Employee.City,
			Address:     order.Employee.Address,
		}
		var empresult *mongo.InsertOneResult
		empresult, err = collection.InsertOne(sc, employee)
		if err != nil {
			sess.AbortTransaction(context.TODO())
			return fmt.Errorf("Could not insert record to employee table : %+v", err)
		}

		collection = helpers.Client.Database("northwind").Collection("customers")
		customer := Customer{
			ContactName: order.Customer.ContactName,
			City:        order.Customer.City,
			Address:     order.Customer.Address,
		}
		var custresult *mongo.InsertOneResult
		custresult, err = collection.InsertOne(sc, customer)
		if err != nil {
			sess.AbortTransaction(context.TODO())
			return fmt.Errorf("Could not insert record to customer table: %+v", err)
		}
		cart := struct {
			OrderDate    time.Time
			ShippedDate  time.Time
			ShipName     string
			ShipAddress  string
			EmployeeInfo interface{}
			CustomerInfo interface{}
			OrderDetails []OrderDetails
		}{
			OrderDate:    time.Now(),
			ShippedDate:  time.Now(),
			ShipName:     order.ShipName,
			ShipAddress:  order.ShipAddress,
			EmployeeInfo: empresult.InsertedID,
			CustomerInfo: custresult.InsertedID,
			OrderDetails: order.OrderDetailsList,
		}
		collection = helpers.Client.Database("northwind").Collection("orders")
		var orderresult *mongo.InsertOneResult
		orderresult, err = collection.InsertOne(sc, cart)

		if err != nil {
			sess.AbortTransaction(context.TODO())
			return fmt.Errorf("Could not insert record to order table : %+v", err)
		}

		sess.CommitTransaction(context.TODO())
		fmt.Printf("Order Document Inserted with _id : %v", orderresult.InsertedID)
		return nil
	})
	sess.EndSession(context.TODO())
	return nil
}
