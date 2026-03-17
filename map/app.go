package main

import "fmt"

// First one is the key type and the 2nd one is the value type of that key

func main() {
	countryCode := map[string]int{
		"India": 91,
		"USA":   1,
		"UK":    44,
	}
	fmt.Println(countryCode)
	fmt.Printf("Country code of India is %v\n", countryCode["India"])

	// Adding new key-value pair to the map
	countryCode["Australia"] = 61
	fmt.Println(countryCode)

	// Deleting a key-value pair from the map
	delete(countryCode, "UK")
	fmt.Println(countryCode)
	
}
