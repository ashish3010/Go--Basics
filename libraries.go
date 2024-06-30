package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := "hello world"

	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.Contains(str, "hello"))
	fmt.Println(strings.ReplaceAll(str, "hello", "hi"))
	fmt.Println(strings.Index(str, "w"))
	fmt.Println(strings.Split(str, ""))
	arr := strings.Split(str, "")
	fmt.Println(strings.Join(arr, ""))

	fmt.Println("original string: ", str)

	ages := []int{23, 55, 30, 40, 5, 10}
	names := []string{"Ashish", "Aashish", "Abhay", "Abdul"}
	sort.Ints(ages)     //sort array
	sort.Strings(names) //sort array

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)
	fmt.Println(ages)
}
