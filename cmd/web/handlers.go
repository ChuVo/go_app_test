package main

import (
	"fmt"
	"go_app_test/pkg/models"
	"net/http"
	"strconv"
)

//TODO Change to map
var usersList []models.User

var DB = make(map[int]models.User)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // Использование помощника notFound()
		return
	}

	w.Write([]byte("Main page!!!"))
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	var currentUser models.User

	if err != nil || id < 1 {
		app.notFound(w) // Использование помощника notFound()
		return
	} else {
		for _, user := range usersList {
			if user.ID == id {
				currentUser.ID = user.ID
				currentUser.UserName = user.UserName
				currentUser.UserEmail = user.UserEmail
				currentUser.Password = user.Password
				currentUser.Admin = user.Admin
				currentUser.Delete = user.Delete
				break
			}
		}

		if currentUser.ID != 0 {
			fmt.Println(currentUser)
		} else {
			fmt.Println("Пользователя с таким ID не существует")
		}
	}

	fmt.Fprintf(w, "User with ID %d...", id)
}

//Register
func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	var randomId = app.makeRandomId()

	user.ID = randomId
	user.UserEmail = "user@email.ru"
	user.UserName = "User Name"
	user.Password = "Password"
	user.Admin = true
	user.Delete = false
	usersList = append(usersList, user)

	fmt.Println("Create new user", user)
	fmt.Println("New list", usersList)
	w.Write([]byte("Добавление нового пользователя"))

	//DB[randomId] = user
	//
	//fmt.Println("Create new user", user)
	//fmt.Println()
	//fmt.Println(DB)
}

func (app *application) getUsersList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users" {
		app.notFound(w)
		return
	}

	if usersList != nil || len(usersList) <= 0 {
		fmt.Println("Users do not exist yet!!!")
	} else {
		fmt.Println("All users", usersList)
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
