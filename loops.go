package main

import "fmt"

func loops() {
	i := 0
	for i < 5 {
		fmt.Println("value of i: ", i)
		i++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i: ", i)
	}

	names := []string{"John", "Jane", "Alex", "Adam"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// range in slice
	for index, value := range names {
		fmt.Printf("at index %v is %v \n", index, value)
	}

	// if u don't want to print index
	for _, value := range names {
		fmt.Printf("value: %v \n", value)
	}

}
