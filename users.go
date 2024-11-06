package main

type User struct {
	Email    string
	Password string
}

var users []User

func addUser(email, password string) {
	users = append(users, User{Email: email, Password: password})
}
