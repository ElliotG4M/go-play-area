package dataTypes

import "fmt"

func variables() {
	// Verbose declaration - Integer
	var i int
	i = 42
	fmt.Println(i)

	// Inline declaration - Float (32 or 64)
	var pi float32 = 3.14
	fmt.Println(pi)

	// Implicit declaration - String
	name := "Elliot Parker"
	fmt.Println(name)

	// Booleans
	boolVar := true
	fmt.Println(boolVar)

	// Declare the complex number 3+4i
	complexNum := complex(3, 4)
	fmt.Println(complexNum)
}

func pointers() {
	// Declare a pointer that points at a string. new(string) required to initialise the pointer
	var firstName *string = new(string)
	// This will print the pointer address e.g. 0x40c128
	fmt.Println(firstName)
	// This will dereference the pointer and save data to the address the pointer points at
	*firstName = "Elliot"
	// This will print 'Elliot'
	fmt.Println(*firstName)

	lastName := "Parker"
	// This will assign a pointer to the lastName variable. & denotes the address of, so here it's the address of lastName
	ptr := &lastName
	// This will print 0x49c128 Parker
	fmt.Println(ptr, *ptr)
	// Update the value of the data but not the location
	lastName = "Burr"
	// This will print 0x49c128 Burr
	fmt.Println(ptr, *ptr)
}

func constants() {
	// Untyped constant
	const untyped = 3
	fmt.Println(untyped + 3)   // Will interpret as an int
	fmt.Println(untyped + 1.2) // Will interpret as a float

	// Typed constant
	const typed int = 3
	//fmt.Println(typed + 1.2) would error as typed is an int
	fmt.Println(float32(typed) + 1.2) // cast typed to a float first

}

// Declare constants globally
const (
	first   = 1
	second  = "second"
	counter = iota
)

// iota constants have the value of the order they're declared. I.e. this one prints 2 as first would be 0, second would be 1, counter is 3rd - so has value 2
func iotaTest() {
	fmt.Println(counter) // prints 2
}

// iota resets in each new constant block
const (
	iotaZero = iota + 6  // 0 + 6
	iotaOne  = 2 << iota // 1 bit shited twice - 4
	iotaTwo              // Inherits the iota function above, so 2 bit shifted twice - 8
)

func iotaTestTwo() {
	fmt.Println(iotaZero, iotaOne, iotaTwo) // Prints 6 4 8
}
