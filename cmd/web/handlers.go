package main

import (
	"fmt"
	"go_app_test/pkg/models"
	"net/http"
	"strconv"
)

var usersList []models.User

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // Использование помощника notFound()
		return
	}

	w.Write([]byte("Main page!!!"))
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Использование помощника notFound()
		return
	}

	var currentUser models.User
	for _, user := range usersList {
		if user.ID == id {
			currentUser.ID = user.ID
			currentUser.UserName = user.UserName
			currentUser.UserEmail = user.UserEmail
			currentUser.Admin = user.Admin
			currentUser.Delete = user.Delete
			break
		}
	}

	fmt.Fprintf(w, "User with ID %d...", id)
	fmt.Println(currentUser)
}

//Register
func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	user.ID = app.makeRandomId()
	user.UserEmail = "user@email.ru"
	user.UserName = "User Name"
	user.Admin = true
	user.Delete = false
	usersList = append(usersList, user)

	fmt.Println("Create new user", user)
	fmt.Println("New list", usersList)
	//w.Write([]byte("Добавление нового пользователя"))
}

func (app *application) getUsersList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users" {
		app.notFound(w)
		return
	}

	if usersList != nil || len(usersList) <= 0 {
		fmt.Println("Users do not exist yet!!!")
	} else {
		fmt.Println(usersList)
	}

	w.Write([]byte("All users"))
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		app.notFound(w)
		return
	}

	w.Write([]byte("Logout"))
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		app.notFound(w)
		return
	}

	w.Write([]byte("Login"))
}

//TODO
//Переделать createUser на регистрацию
//Сделать валидацию па правильность почты, эксклюзивность почты
//Посмотреть библиотеку на генерацию ID
func (app *application) register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		app.notFound(w)
		return
	}

	w.Write([]byte("register"))
}
