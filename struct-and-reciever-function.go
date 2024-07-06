package main

import "fmt"

type bill struct {
	name string
	item map[string]float64
	tip  float64
}

func newBill(name string) bill {
	bill1 := bill{
		name: name,
		item: map[string]float64{},
		tip:  0,
	}

	return bill1
}

// reciever function
func (b bill) formatBill() string {
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

// update tip , since we are updating the original struct, we will use pointers
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item, it will add to original struct
func (b *bill) addItem(name string, price float64) {
	b.item[name] = price
}

func structnReciverFunctions() {
	bill := newBill("John Doe")

	bill.updateTip(28.00)

	bill.addItem("cake", 8.54)
	bill.addItem("pie", 3.64)

	fmt.Println(bill.formatBill())
}
