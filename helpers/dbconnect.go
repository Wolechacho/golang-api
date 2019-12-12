package helpers

import (
	"context"
	"first-api-golang/models/utilitymodel"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

//Client -- mongo client
var Client *mongo.Client
var once sync.Once
var err error

//ConnectToMongoDb -- coonection to the mongodb database
func ConnectToMongoDb(connURI string) *mongo.Client {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(connURI)
		Client, err = mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			log.Fatal(err)
		}
		err = Client.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to MongoDB!")
	})
	return Client
}

//FormatConnectionString - format the representaion of a mongodb conn uri
func FormatConnectionString(filename string) string {

	//Read the file content to bytes
	data := ReadYamlFromFile(filename)
	c := utilitymodel.ConnectionString{}

	//parse yaml config to connection string struct
	err := yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalln("Error Marshalling : ", err)
	}

	var sb strings.Builder
	sb.WriteString("mongodb://")

	if c.UserName != "" || c.Password != "" {
		sb.WriteString(fmt.Sprintf("@%v:%v@", c.UserName, c.Password))
	}
	for _, endpoint := range c.Endpoints {
		sb.WriteString(fmt.Sprintf("%v:%v,", endpoint.IP, endpoint.Port))
	}

	//remove the last occurence of comma (,)
	s := sb.String()
	s = s[:len(s)-1]

	//Clear the StringBuilder
	sb.Reset()
	sb.WriteString(fmt.Sprintf("%v/%v?replicaSet=rs", s, c.DatabaseName))

	return sb.String()
}

//ReadYamlFromFile -- read the file content using filename
func ReadYamlFromFile(filename string) []byte {
	if filename == "" {
		log.Fatalln("FileName in Empty")
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Could not read file : ", err)
	}
	return b
}