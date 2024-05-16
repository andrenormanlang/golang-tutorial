// go build -o myprogram main.go
// ./myprogram

package main

import (
	"fmt"
)

// * is a pointer to the memory address of the value
func updateName(x * string) {
	*x = "wedge"
}

func main () {
	name := "tifa"
	// updateName(name)

	// fmt.Println("memory address of name is:", &name)

	m := &name
	// memory address of the pointer
	// fmt.Println("memory address of m is:", m)
	// fmt.Println("value at memory address m is:", *m)

	// updating the name using the pointer to the memory address of the value
	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

}

/* 
|---name---|----m----|
|  0x001   | 0x002  |
|  tifa   | p0x001  |
|---------|--------|
*/

func updateName(x string) string{
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64)  {
	y["coffee"] = 2.99
}

func main () {
	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	// updating a copy of the original value of the function
	// non-pointer to the original value of the function in memory
	name = updateName(name)
	fmt.Println(name)

	// group B types -> slices, maps, functions
	menu := map[string]float64{
		"pie": 5.95,
		"ice cream": 3.99,
	}
	// pointer to the original value of the function in memory
	// pointer wrapper value
	updateMenu(menu)
	fmt.Println(menu)

}

// Map is a collection of key value pairs
func main() {
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee": 3.55,
		"coffee": 1.99,
		"muffin": 2.45,
		"tea":  1.49,
		"sandwich": 5.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// loop through the map
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		+46769182406: "Mario",
		+46769202305: "Luigi",
		+46769004500: "Peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[+46769182406])

	// update the value
	phonebook[+46769182406] = "Bowser"
	fmt.Println(phonebook)

}

import (
	"fmt"
)

// has to be declared outside of the main function
var score = 99.5

func main() {	
	sayHello("Mario")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}

Function returning multiple values
import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	fn1, sn2 := getInitials("tifa lockhart")
	fmt.Println(fn1, sn2)

	fn3, sn4 := getInitials("cloud strife")
	fmt.Println(fn3, sn4)

	fn3, sn3 := getInitials("yuffie")
	fmt.Println(fn3, sn3)

}

import(
	"fmt"
	"math"
)

// functions in GO
func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// functions in GO
func main() {
	// sayGreeting("Mario")
	// sayGreeting("Luigi")
	// sayBye("Mario")

	cycleNames([]string{"Mario", "Luigi", "Peach"}, sayGreeting)
	cycleNames([]string{"Mario", "Luigi", "Peach"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("Circle 1 is %.3f and Circle 2 is %.3f \n", a1, a2)
}

If statement and else if statement
func main(){

	age := 30

	fmt.Println( age <= 50 )
	fmt.Println( age >= 50 )
	fmt.Println( age == 50 )
	fmt.Println( age != 50 )
	fmt.Println( age == 30 )

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age == 30 {
		fmt.Println("age is 30")
	} else {
		fmt.Println("age is not less than 30")
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}

		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		

		fmt.Printf("the value at index %v is %v \n", index, value)
	}
}

func main(){
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is:", i)
	// }

	// iterate over a collection
	names := []string{"mario", "luigi", "yoshi", "peach"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	
	// similar to foreach loop in javascript
	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	// }

	// if you don't want to use the index
	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		// does not change the original array
		value = "new string"
	}

	fmt.Println(names)

}

go build main.go
./main

import (
	"fmt"
	"sort"
	// "strings"
)

func main(){
	greeting := "Hello there friends!"

	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "th"))
	// fmt.Println(strings.Split(greeting, " "))


	// The original string is not modified
	fmt.Println("Original string value: ", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	// returns 
	index2 := sort.SearchInts(ages, 90)
	fmt.Println(index2)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))
}


func main() {
	// array
	//  array is curly braces and the length of the array 
	// unlike javascript, the length of the array is fixed
	var ages = [3]int{20, 25, 30}

	fmt.Println(ages, len(ages))

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	// slices are dynamic arrays
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges
	// returns  1 and 2
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne , rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
}

func main() {

	age := 35
	name := "Shaun"

	// Print
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line! \n")

	// Println
	fmt.Println("Hello, ninjas!")
	fmt.Println("Goodbye, ninjas!")
	fmt.Println("My age is", age, "and my name is", name)

	// formatted strings %_ = format specifier
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	// %q = double quotes and string types
	fmt.Printf("My age is %q and my name is %q \n", age, name)
	// %T = type of the variable
	fmt.Printf("age is of type %T \n", age)
	// %f = float with decimal points
	fmt.Printf("You scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)

	fmt.Println("The saved string is:", str)
}

func main() {
	// // strings
	// var nameOne string = "mario"
	// // infers the type of variable as string
	// var nameTwo = "luigi"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)
	
	// nameOne = "peach"
	// nameThree = "bowser"
	
	// fmt.Println(nameOne, nameTwo, nameThree)

	// // short declaration without var keyword and type
	// nameFour := "yoshi"

	// fmt.Println(nameFour)


	// integers - floats or int

	// int
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory - check the scope of the variable
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint16 = 256 // no negative numbers for unint

	// float - specify the bit size
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 8888885939898530853.7
	// infers the type of variable as float64
	scoreThree := 1.5


}