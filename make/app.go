package main

import "fmt"

// Make is used to create slices, maps and channels. It returns an initialized (not zeroed) value of type T (not *T).
// Take 3 argument, 1st is type 2nd one is the length and 3rd one is the capacity of slice. If the capacity is not provided,
//  it will be set to the length of the slice.
// Once the length of the slice reaches its capacity, it will automatically increase the capacity by doubling it.

func main() {
	const intialSize = 5
	const capacity = 10
	userNames := make([]string, intialSize, capacity)
	userNames[0] = "John"
	userNames[1] = "Jane"
	userNames[2] = "Doe"
	userNames[3] = "Smith"
	userNames[4] = "Johnson"

	// userNames[5] = "Williams" // This will cause an error because the length of the slice is 5 and we are trying to access the index 5 which is out of range.

	// To add new element to the slice, we can use the append function which will automatically increase the length of the slice and also increase the capacity if needed.
	userNames = append(userNames, "Williams")
	userNames = append(userNames, "Williams")
	userNames = append(userNames, "Williams")
	userNames = append(userNames, "Williams")
	userNames = append(userNames, "Williams")
	userNames = append(userNames, "Williams")
	userNames = append(userNames, "Williams")
	userNames = append(userNames, "Williams")
	userNames = append(userNames, "Williams")
	userNames = append(userNames, "Williams")
	userNames = append(userNames, "Williams")
	userNames = append(userNames, "Williams")
	userNames = append(userNames, "Williams")
	fmt.Println(userNames)

}
