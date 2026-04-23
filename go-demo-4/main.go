package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc account) outpudPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

var letterRune = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ1234567890-*!")

func main() {
	fmt.Println(generatePassword(10))
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount := account{
		login:    login,
		password: password,
		url:      url,
	}
	myAccount.outpudPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

func generatePassword(n int) string {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRune[rand.IntN(len(letterRune))]
	}
	return string(result)
}
