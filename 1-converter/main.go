package main

import (
	"fmt"
)

func main() {
	const USDToEUR = 0.9
	const USDToRUB = 80
	EURToRUB := USDToRUB / USDToEUR

	//	U.R = USD * RUB
	fmt.Println(EURToRUB)
}
