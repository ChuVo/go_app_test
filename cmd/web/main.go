package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Get Start!")

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/user", showUser)
	mux.HandleFunc("/user/create", createUser)

	log.Println("Запуск сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
