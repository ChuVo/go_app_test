package main

import (
	"fmt"
	"go_app_test/pkg/models"
)

var dataBase = make(map[string]models.User)

func addItem() {

	dataBase = make(map[string]models.User)
	dataBase["ID 1212"] = models.User{
		1123,
		"Joe",
		"password",
		"1212",
		true,
		true,
	}

	fmt.Println(dataBase)
}
