package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"first-api-golang/models"
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
