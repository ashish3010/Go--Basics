package main

import "fmt"

func main() {
	// arrays
	var ages [3]int = [3]int{25, 38, 52}
	names := [4]string{"John", "Doe", "Jane", "Doey"}

	fmt.Println(ages, len(ages), ages[1])
	fmt.Println(names, len(names), names[1])

	// slices
	scores := []int{100, 85, 21}
	fmt.Println(scores, len(scores))

	scores = append(scores, 92) //append returns new array
	fmt.Println(scores, len(scores))

	// slice ranges
	range1 := scores[1:3] // will slice score array from index 1(included) to 3(excluded)
	range2 := scores[1:]  // will slice score array from index 1(included) to last element
	range3 := scores[:3]  // will slice score array from index 0(included) to 3(excluded)

	range1 = append(range1, 32) //will add 31 in range array

	fmt.Println(range1, range2, range3)
}
