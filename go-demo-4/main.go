package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByURL,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите вариант",
}

func main() {
	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		variant := promptData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
}

func findAccountByURL(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccount(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин для поиска")
	accounts := vault.FindAccount(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунты не найдены")
	}
	for _, account := range *accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	isDelited := vault.DeleteAccountByULR(url)
	if isDelited {
		color.Green("Аккаунт удален")
	} else {
		output.PrintError("Не найдено")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или логин")
		return
	}
	vault.AddAccount(*myAccount)
	color.Green("Запись успешна")
}

func promptData(prompt ...string) string {
	for i, v := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", v)
		} else {
			fmt.Println(v)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
