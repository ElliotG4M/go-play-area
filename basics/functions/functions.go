package functions

import (
	"errors"
	"fmt"
)

// Standard function definition
func StartWebServer(port int, retries int) (int, error) {
	fmt.Println("Starting server...")
	// Define an anonymous inline function. Parameter definition goes in the normal location, passing in the parameter happens at the end of the function so...
	func(port int) {
		fmt.Println("Using port", port)
	}(port) //... <-- parameter is passed in here

	// Start server
	fmt.Println("Startup complete!")
	return port, errors.New("Something went wrong")
}
