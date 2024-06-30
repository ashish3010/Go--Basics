package main

import "fmt"

func main() {

	// decalaration

	// string
	var name = "8734"
	var name1 string = "John"
	name2 := "John"
	var x string

	fmt.Println(name, name1, name2, x)

	// https://pkg.go.dev/builtin   //all ranges

	// integer
	var num1 int = 21   //no range
	var num2 = 21       //no range
	var num3 int8 = 127 //ranged
	var num4 uint = 23  //only accept positive
	num5 := 34

	fmt.Println(num1, num2, num3, num4, num5)

	// float
	var f1 float32 = 21.00
	var f2 float64 = 217.00843673846436445676543

	fmt.Println(f1, f2)
}
