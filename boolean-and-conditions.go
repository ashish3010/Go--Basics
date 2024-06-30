package main

import "fmt"

func main() {
	age := 233

	fmt.Println(age < 30)
	fmt.Println(age > 25)
	fmt.Println(age <= 25)
	fmt.Println(age >= 25)
	fmt.Println(age == 23)

	if age < 20 {
		fmt.Println("age is less thant 20")
	} else if age == 23 {
		fmt.Println("age is 23")
	} else {
		fmt.Println("age is greater than 23")
	}

	names := []string{"John", "Doe", "Jane", "Doey"}
	for index, value := range names {
		if index == 1 {
			fmt.Printf("continue at index %v \n", index)
			continue
		}

		if index >= 2 {
			fmt.Printf("break at index %v \n", index)
			break
		}

		fmt.Printf("index: %v, value: %v \n", index, value)
	}
}
