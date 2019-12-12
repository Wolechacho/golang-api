package controllers

import (
	"encoding/json"
	"first-api-golang/models"
	"fmt"
	"io/ioutil"
	"net/http"
)

//OrderController routes
type OrderController struct {
}

//Save -- api route to save order
func (o *OrderController) Save(w http.ResponseWriter, r *http.Request) {
	var responseCode string
	var message string

	reqBody, err := ioutil.ReadAll(r.Body)
	if err == nil {
		s := string(reqBody)
		fmt.Println("string representation of the converted bytes : ", s)

		var order models.Order
		err = json.Unmarshal(reqBody, &order)
		if err == nil {
			err = order.SaveOrder(&order)
			if err == nil {
				responseCode = "00"
				message = "Order was successfuly saved"
			} else {
				responseCode = "10"
				message = fmt.Sprintf("Order was not successfuly saved : %v", err)
			}
		}else {
			responseCode = "10"
			message = fmt.Sprintf("could not unmarshal the request body : %v", err)
		}
	}else {
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
