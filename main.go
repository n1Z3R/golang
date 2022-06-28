package main

import (
	"fmt"
	"golang-module/user"
)

func main() {
	var name, email, password string

	fmt.Println("Введите имя:")
	fmt.Scanln(&name)
	fmt.Println("Введите почту:")
	fmt.Scanln(&email)
	fmt.Println("Введите пароль:")
	fmt.Scanln(&password)

	user, status := user.SetUser(name, email, password)
	if !status {
		fmt.Println("Почта или пароль (минимум 8 символов) введены неверно")
	} else {
		fmt.Println("Успешная регистрация!")
		fmt.Printf("%+v\n", user)
	}

}
