package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// bill structure
type Bill struct {
	name string
	item map[string]float64
	tip  float64
}

// creating a bill
func newEmptyBill(name string) Bill {
	bill1 := Bill{
		name: name,
		item: map[string]float64{},
		tip:  0,
	}
	return bill1
}

// adding item to original bill
func (b *Bill) addItemToBill(name string, price float64) {
	b.item[name] = price
}

// taking input from user
func getInfo(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
}

// creating a bill from user input
func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInfo("Create a new bill: ", reader)

	b := newEmptyBill(name)
	fmt.Println("Bill created - ", name)

	return b
}

// formatting a bill
func (b Bill) formatBill() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	for k, v := range b.item {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	return fs
}

// adding tip to bill
func (b *Bill) updateTip(tip float64) {
	b.tip = tip
}

// saving a bill in bills folder
func (b *Bill) save() {
	data := []byte(b.formatBill())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved to file")
}

// taking users options using switch statement
func promptOptions(b Bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInfo("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInfo("Item name: ", reader)
		price, _ := getInfo("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.addItemToBill(name, p)
		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "s":
		b.save()
	case "t":
		tip, _ := getInfo("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip given - ", tip)
		promptOptions(b)
	default:
		fmt.Println("choose a valid option...")
		promptOptions(b)
	}
}

func userInputAndSavingBillToBillsFolder() {
	bill := createBill()

	promptOptions(bill)
}
