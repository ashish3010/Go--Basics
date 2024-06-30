package main

import (
	"fmt"
	"math"
	"strings"
)

func greet(n string) {
	fmt.Printf("hello %v \n", n)
}

func bye(n string) {
	fmt.Printf("bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// return multiple value
func getInitials(n string) (string, string) {
	str := strings.ToUpper(n)
	names := strings.Split(str, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func functions() {
	greet("John")
	bye("John")

	cycleNames([]string{"John", "Jane", "Alex"}, greet)
	cycleNames([]string{"John", "Jane", "Alex"}, bye)

	circle1 := circleArea(25)
	circle2 := circleArea(12.45)

	fmt.Printf("area of circle 1 is %0.2f \n", circle1)
	fmt.Printf("area of circle 2 is %0.3f \n", circle2)

	fn1, sn1 := getInitials("John Doe")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("John")
	fmt.Println(fn2, sn2)
}
