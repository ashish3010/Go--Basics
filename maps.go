package main

import "fmt"

func maps() {
	menu := map[string]float32{
		"soup":   8.32,
		"milk":   2.45,
		"bread":  5.12,
		"coffee": 6,
	}

	fmt.Println(menu)
	fmt.Println(menu["milk"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phonebook := map[int]string{
		9876543210: "John",
		1234567890: "Jane",
		2354234453: "Janardhan",
		8478374844: "Jordan",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[2354234453])

	phonebook[2354234453] = "JD"
	phonebook[8478374844] = "JJ"

	fmt.Println(phonebook)
}
