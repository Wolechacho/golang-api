package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Category struct {
	CategoryName string
	Description  string
}

type Product struct {
	ProductName  string
	UnitPrice    float64
	UnitInStock  int
	Discontinued bool
	CategoryInfo primitive.ObjectID
}
type Order struct {
	OrderDate    time.Time
	ShippedDate  time.Time
	ShipName     string
	ShipAddress  string
	OrderDetails []OrderDetails
	EmployeeInfo interface{}
	CustomerInfo interface{}
}

type OrderDetails struct {
	UnitPrice   float64
	Quantity    int
	Discount    float64
	ProductInfo primitive.ObjectID
}

type Employee struct {
	ContactName string
	Address     string
	City        string
}

type Customer struct {
	ContactName string
	Address     string
	City        string
}

var client *mongo.Client
var err error

func main() {

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", homeLink)
	// log.Fatal(http.ListenAndServe(":3030", router))
	connectToMongoDb()
	saveOrder()
	//products := getProducts()
	// product := getProductByID("5ddc700409c0682e5fe95a7c")
	// fmt.Printf("Single Product found : %+v\n", product)

	//deleteProductByID("5ddc700409c0682e5fe95a7c")

	//fmt.Printf("Products found : %+v\n", products)
	//saveProduct()
	//saveCategory()
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

//connec to the mongodb database
func connectToMongoDb() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017,localhost:27018,localhost:27019/northwind?replicaSet=rs")
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func saveCategory() {
	collection := client.Database("northwind").Collection("categories")

	newcategory := Category{CategoryName: "Kitchen Utensils", Description: "Used for house chores"}

	insertResult, err := collection.InsertOne(context.TODO(), newcategory)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Category Document : ", insertResult.InsertedID)
}

func saveProduct() {
	collection := client.Database("northwind").Collection("products")

	categoryid, err := primitive.ObjectIDFromHex("5ddae874a4c9080554c1b4e3")
	if err != nil {
		log.Fatalln("Could not convert hex number to ObjectId")
	}
	newproduct := Product{
		ProductName:  "Yeezy",
		UnitPrice:    200.0,
		UnitInStock:  4,
		Discontinued: false,
		CategoryInfo: categoryid,
	}

	insertResult, err := collection.InsertOne(context.TODO(), newproduct)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Product Document: ", insertResult.InsertedID)
}

func getProducts() []Product {
	var fetchedProducts []Product
	collection := client.Database("northwind").Collection("products")

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

func getProductByID(id string) Product {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	var product Product
	collection := client.Database("northwind").Collection("products")
	err = collection.FindOne(context.TODO(), bson.M{"_id": oid}, options.FindOne()).Decode(&product)
	if err != nil {
		log.Fatal(err)
	}
	return product
}

func deleteProductByID(id string) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("northwind").Collection("products")
	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": oid}, options.Delete())
	if err != nil {
		log.Fatal(err)
	}
	if result.DeletedCount != 1 {
		log.Fatal(err)
	}
	fmt.Printf("Number of product deleted  : %d\n", result.DeletedCount)
}

func saveOrder() {
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
