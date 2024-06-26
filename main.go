package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// helper function to get input from the user
func getInput(prompt string, r *bufio.Reader)(string, error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill(name string) bill {
	// reads the input from the user in the terminal
	reader := bufio.NewReader(os.Stdin)

	// fmt.Println("Create a new bill name:")
	// // reads the input after pressed enter
	// name, _ = reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ = getInput("Create a new bill name:", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOptions (b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip):", reader)
	// fmt.Println(opt)

	switch opt {
		case "a":
			name, _ := getInput("Item name:", reader)
			price, _ := getInput("Item price:", reader)

			// convert the price to a float
			p, err := strconv.ParseFloat(price, 64)
			if err != nil {
				fmt.Println("The price must be a number")
				promptOptions(b)
			}
			b.addItem(name, p)

			fmt.Println(name, price)

			fmt.Println("Item added -", name, price)
			promptOptions(b)
		case "t":
			tip, _ := getInput("Enter tip amount ($):", reader)
			t, err := strconv.ParseFloat(tip, 64)
			if err != nil {
				fmt.Println("The tip must be a number")
				promptOptions(b)
			}
			b.updateTip(t)

			fmt.Println("Tip added:", tip)
			promptOptions(b)
		case "s":
			b.save()
			fmt.Println("You saved the bill", b)
		default:
			fmt.Println("That is not a valid option...")
			promptOptions(b)
	}
}

func main() {
	mybill := createBill("Pixie's bill")

	promptOptions(mybill)

	// fmt.Println(mybill)

	// mybill := newBill("Pixie's bill")

	// mybill.addItem("Onion soup", 4.50)
	// mybill.addItem("veg pie", 8.95)
	// mybill.addItem("tofu curry", 10.95)
	// mybill.addItem("pasta", 9.95)


	// // adds items to the bill
	// mybill.updateTip(10)

	// // prints the bill with the formatted string
	// fmt.Println(mybill.format())
}