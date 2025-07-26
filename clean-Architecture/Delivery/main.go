package main

import "task-management/Delivery/router"

func main() {
	databaseClient := router.InitDB("mongodb://localhost:27017")
	// creates all the controllers and start connection with mongoDB database
	router.SetUp(databaseClient)
	defer router.CloseDB(databaseClient)

}
