package main

import (
	"log"
	"net/http"

	"first-api-golang/helpers"
	"first-api-golang/routers"
)

func main() {
	helpers.ConnectToMongoDb()

	router := routers.RegisterRoutes()
	err := http.ListenAndServe(":3031", router)
	if err != nil {
		log.Fatal(err)
	}
}
