package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"first-api-golang/models"

	"github.com/gorilla/mux"
)

//ProductController routes
type ProductController struct {
}

//Save -- api route to insert product
func (pc *ProductController) Save(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("could not read the request body")
	}
	s := string(reqBody)
	fmt.Println("string representation of the converted bytes : ", s)
	var product models.Product
	err = json.Unmarshal(reqBody, &product)
	if err != nil {
		log.Fatalln("could not umarshall the request body : ", err)
	}
	product.SaveProduct(&product)
}

//List -- api route to get all products
func (pc *ProductController) List(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	list := product.GetProducts()
	fmt.Println(list)
}

//Get -- api route to get a product by id
func (pc *ProductController) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product := models.Product{}
	p := product.GetProductByID(params["id"])
	fmt.Println(p)
}

//Delete -- Removes a product document from the DB
func (pc *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product := models.Product{}
	product.DeleteProductByID(params["id"])
}
