/*
 When the value is passed to the function, a copy is being created and later it will be garbage collected.
 For large complex data structures, this can be inefficient. Instead of passing the value,
 we can pass a pointer to the value.


 Direct value can be mutated using pointers.
*/

package main

import "fmt"

func main() {

	age := 100

	// Print the address of the age

	fmt.Printf("Address of age : %v\n", &age)

	updatedAge := getUpdatedAge(&age)

	fmt.Printf("Updated age : %v\n", updatedAge)

	fmt.Printf("Age before mutation : %v\n", age)

	mutateAge(&age)
	fmt.Printf("Age after mutation : %v\n", age)

}

// Getting the memory address instead of creating the new variable
func getUpdatedAge(age *int) int {
	// Accessing the value on that memory address
	return *age - 14
}

// This will directly mutate the value of age without creating a new variable
func mutateAge(age *int) {
	*age = *age - 14
}
