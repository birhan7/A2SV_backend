package main

import (
	"fmt"
	"task3/library_management/controllers"
	"task3/library_management/models"
	"task3/library_management/services"
)

func main() {
	fmt.Println("...Library Management System is running...")
	lib := &services.Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
	controllers.Menu(lib)
}
