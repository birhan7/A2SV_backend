package main

import (
	"task-manager-api/router"
)

func main() {
	router := router.SetupRouter()
	router.Run(":9000")
}
