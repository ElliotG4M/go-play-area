package dataTypes

import "fmt"

func arrays() {
	// Verbose
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1) // [1 2 3]

	// Inline
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2) // [1 2 3]
}

// slices are built on arrays
func slices() {
	arr := [3]int{1, 2, 3}

	// arr[:] tells the slice to be created with all of arr's contents
	slice := arr[:]

	fmt.Println(arr, slice) // [1 2 3] [1 2 3]

	// The slice and array are linked as follows:
	arr[0] = 12
	slice[1] = 99
	fmt.Println(arr, slice) // [12 99 3] [12 99 3]

	// No size or original array specified - compiler will generate underlying array
	slice2 := []int{1, 2, 3}
	// Add new items to the slice. Slices are not fixed size - Go will automatically regenerate the underlying array.
	// For very large slices, it can be worth using some in built functions to reserve the memory for the array
	slice2 = append(slice2, 4, 5)
	fmt.Println(slice2) // [1 2 3 4 5]

	// Generate slice from a subset of another array or slice
	slice3 := slice[1:]                 // gets the slice from element 1 onwards 						[12 99 3] => [99 3]
	slice4 := slice[:2]                 // gets the slice up to but not including element 2 			[12 99 3] => [12 99]
	slice5 := slice[1:2]                // gets the slice from element 1 to element 2 (non inclusive) 	[12 99 3] => [99]
	fmt.Println(slice3, slice4, slice5) // [99 3] [12 99] [99]

}

func maps() {
	m := map[string]int{"foo": 42} // map[keyType]valueType{json like definition}
	fmt.Println(m["foo"])          // 42

	// Update map contents
	m["foo"] = 128

	// remove an item from a map
	delete(m, "foo")
}

func structs() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	// Initialised with 1 element with default values
	var u user
	fmt.Println(u) // {0  } (i.e. {0 '' ''})

	u.ID = 1
	u.FirstName = "Elliot"
	u.LastName = "Parker"
	fmt.Println(u)           // {1 Elliot Parker}
	fmt.Println(u.FirstName) // Elliot

	u2 := user{
		ID:        1,
		FirstName: "Elliot",
		LastName:  "Parker", // end with an extra comma or Go will think this is the end of a line
	}

	fmt.Println(u2)
}
