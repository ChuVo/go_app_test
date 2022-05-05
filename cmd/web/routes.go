package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/user", app.showUser)
	mux.HandleFunc("/user/create", app.createUser)
	mux.HandleFunc("/users", app.showUsersList)
	
	return mux
}
