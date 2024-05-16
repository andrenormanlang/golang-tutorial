// go build -o myprogram main.go
// ./myprogram

// Struct -> collection of fields and methods
package main

import (
	"fmt"
	"os"
)

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// make a new bill
func newBill(name string) bill {
	// below is based on the zero value of the struct
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}

	return b
}

// format the bill
// receiving the bill as a receiver function
// for structs pointers are automaticcally direferenced
func (b *bill) format() string {
	// bill breakdown - fs is the string that will be returned
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs +=  fmt.Sprintf("%-25v ...$%v \n", k+ ":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total + b.tip)

	return fs
}

// update the tip
// whenever updating a value in a struct, use a pointer receiver *
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save the bill
func (b *bill) save() {
	data := []byte(b.format())
	
	// os.WriteFile("bills/" + b.name + ".json", data, 0644)
	err := os.WriteFile("bills/" + b.name + ".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}