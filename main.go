package main

import (
	"log"

	"go.mod/app"
)

func main() {
	// TASK 1
	log.Println(app.Multiply(7.8, 2))

	// TASK 2
	if isValid, err := app.IsInteger("hello"); err != nil {
		log.Println(isValid, err.Error())
	}

	// TASK 3
	log.Println(app.IntToDayName(7))

	// TASK 4
	log.Println(app.DayNameToInt("Senin"))

	// TASK 5
	var kosong string
	app.IsHello(&kosong)
	log.Println(kosong)
}
