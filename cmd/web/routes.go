package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", app.home)
	router.HandleFunc("/user", app.getUser)
	router.HandleFunc("/user/create", app.createUser)
	router.HandleFunc("/users", app.getUsersList)
	router.HandleFunc("/logout", app.logout)
	router.HandleFunc("/login", app.login)
	router.HandleFunc("/register", app.register)

	return router
}
