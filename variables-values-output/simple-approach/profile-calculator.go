/*
	var and const can be defined as a variable and constant cannot be reassigned while var can be reassigned.

	In order to get the user input using scan you need to pass the variable point with & to store the value
	get by the user.

	short cut of assigning values is like below
	value := 100

*/

package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scanln(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scanln(&expenses)

	fmt.Print("Enter tax rate (as a decimal): ")
	fmt.Scanln(&taxRate)

	earningBeforeTax := revenue - expenses

	earningAfterTax := earningBeforeTax - (earningBeforeTax * (taxRate / 100))

	ratio := earningBeforeTax / earningAfterTax

	fmt.Printf("Earning before tax = %f \n", earningBeforeTax)

	fmt.Printf("Profit = %f \n", earningAfterTax)

	fmt.Printf("Ratio = %f", ratio)
}
