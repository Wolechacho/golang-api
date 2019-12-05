package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"first-api-golang/models"

	"github.com/gorilla/mux"
)

//ProductController routes
type ProductController struct {
}

//Save -- api route to insert product
func (pc *ProductController) Save(w http.ResponseWriter, r *http.Request) {
	var responseCode string
	var message string

	reqBody, err := ioutil.ReadAll(r.Body)
	if err == nil {
		s := string(reqBody)
		fmt.Println("string representation of the converted bytes : ", s)

		var product models.Product
		err = json.Unmarshal(reqBody, &product)
		if err == nil {
			err = product.SaveProduct(&product)
			if err == nil {
				responseCode = "00"
				message = "Product was successfuly saved"
			} else {
				responseCode = "10"
				message = fmt.Sprintf("Product was not successfuly saved : %v", err)
			}
		} else {
			responseCode = "10"
			message = fmt.Sprintf("could not unmarshal the request body : %v", err)
		}
	} else {
		responseCode = "10"
		message = fmt.Sprintf("could not read the request body %v", err)
	}
	resp := struct {
		ResponseCode string
		Message      string
	}{
		ResponseCode: responseCode,
		Message:      message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}

//List -- api route to get all products
func (pc *ProductController) List(w http.ResponseWriter, r *http.Request) {
	var responseCode string
	product := models.Product{}
	list, err := product.GetProducts()

	if err != nil {
		responseCode = "10"
		//this needs to be changed to nil
		list = []models.Product{}
	} else {
		responseCode = "00"
	}

	resp := struct {
		ResponseCode string
		Products     []models.Product
	}{
		ResponseCode: responseCode,
		Products:     list,
	}
	fmt.Println(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

//Get -- api route to get a product by id
func (pc *ProductController) Get(w http.ResponseWriter, r *http.Request) {
	var responseCode string
	params := mux.Vars(r)
	product := models.Product{}
	p, err := product.GetProductByID(params["id"])

	if err != nil {
		responseCode = "10"
	} else {
		responseCode = "00"
	}

	resp := struct {
		ResponseCode string
		Product      interface{}
	}{
		ResponseCode: responseCode,
		Product:      p,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

//Delete -- Removes a product document from the DB
func (pc *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	var responseCode string
	var message string
	params := mux.Vars(r)
	product := models.Product{}
	err := product.DeleteProductByID(params["id"])
	if err != nil {
		responseCode = "10"
		message = "Product not deleted"
	} else {
		responseCode = "00"
		message = "Product Deleted"
	}

	resp := struct {
		ResponseCode string
		Message      string
	}{
		ResponseCode: responseCode,
		Message:      message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
