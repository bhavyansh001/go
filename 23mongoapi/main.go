package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/bhavyansh001/mongoapi/route"
)

// https://github.com/mongodb/mongo-go-driver
//  go get go.mongodb.org/mongo-driver/v2/mongo
// get instance of mongo db: https://www.mongodb.com/products/platform/atlas-database || run locally

// root dir has just one go file as per go

// go build main.go


func main() {
	fmt.Println("MongoDB with Go!")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4001", r))
	fmt.Println("Listening at port 4001")
}
