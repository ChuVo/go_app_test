package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Get Start!")

	// Создаем новый флаг командной строки, значение по умолчанию: ":4000".
	// Добавляем небольшую справку, объясняющая, что содержит данный флаг.
	// Значение флага будет сохранено в переменной addr.
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/user", showUser)
	mux.HandleFunc("/user/create", createUser)
	mux.HandleFunc("/users", showUsersList)

	log.Printf("Запуск сервера на %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
