// go build -o myprogram main.go
// ./myprogram

// Struct -> collection of fields and methods
package main

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