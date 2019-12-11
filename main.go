package main

import (
	"context"
	"first-api-golang/helpers"
	"first-api-golang/routers"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	//ServerPort port ex : 3030
	ServerPort = 3030

	//EnvFile ex :env.yml
	EnvFile = "env.yml"
)

func main() {
	connString := helpers.FormatConnectionString(EnvFile)
	mongoClient := helpers.ConnectToMongoDb(connString)

	router := routers.RegisterRoutes()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", ServerPort),
		Handler: router,
	}
	//create channel to the start the server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("API Server started")
	<-done

	//Receive Shutdown signal
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed : %+v", err)
	}
	fmt.Println("API Server shutdown Properly")

	err := mongoClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("Unable to shutdown DB %+v", err)
	}

	fmt.Println("DB Disconnected !!!")
}
