package main

import "fmt"

func formatSpecifier() {

	fname := "John"
	lname := "Doe"
	age := 23

	// Print prints all values in single line
	// if u want to break line add  \n
	fmt.Print("hello")
	fmt.Print("world \n", fname)

	// Println print value in next line
	fmt.Println("hello")
	fmt.Println("world")

	// Printf (formatted string) %_ = format specifier
	fmt.Printf("your age is %v and your name is %v \n", age, fname)  //print value
	fmt.Printf("your name is %q %q \n", fname, lname)                //print value in quotes but only works for strings
	fmt.Printf("age is of type %T \n", age)                          // print type of age
	fmt.Printf("your average score is %f \n", 28.28)                 // print float value
	fmt.Printf("your average with only 1 dcimal if %0.1f \n", 28.28) //print only 1 decimal

	// Sprintf (Save printf)
	str := fmt.Sprintf("your age is %v and your fname is %v and your lname is %v \n", age, fname, lname)
	fmt.Println(str)
}
