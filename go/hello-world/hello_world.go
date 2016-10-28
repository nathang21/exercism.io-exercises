package hello

import "fmt"

const testVersion = 2

// HelloWorld Prints Hello World to the screen
func HelloWorld(name string) string {
	if name != "" {
		return fmt.Sprintf("Hello, %v!", name)
	}
	return fmt.Sprintf("Hello, World!")
}
