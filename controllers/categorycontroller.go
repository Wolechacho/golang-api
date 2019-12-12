package controllers

import (
	"encoding/json"
	"first-api-golang/models"
	"fmt"
	"io/ioutil"
	"net/http"
)

//CategoryController routes
type CategoryController struct {
}

//Save -- api route to insert category
func (c *CategoryController) Save(w http.ResponseWriter, r *http.Request) {
	var responseCode string
	var message string

	reqBody, err := ioutil.ReadAll(r.Body)
	if err == nil {
		var category models.Category
		err = json.Unmarshal(reqBody, &category)

		if err == nil {
			err = category.SaveCategory(&category)

			if err == nil {
				responseCode = "00"
				message = fmt.Sprintf("Category Saved Successfully")
			} else {
				responseCode = "10"
				message = err.Error()
			}

		} else {
			responseCode = "10"
			message = fmt.Sprintf("could not unmarshal the request body : %+v", err)
		}
	} else {
		responseCode = "10"
		message = fmt.Sprintf("could not read the request body %+v", err)
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
