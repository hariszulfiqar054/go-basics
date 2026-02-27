package main

import "fmt"

func main() {

	revenue := getInput("revenue: ")
	expenses := getInput("expenses: ")
	taxRate := getInput("tax rate")

	earningBeforeTax, earningAfterTax := getEarning(revenue, expenses, taxRate)

	ratio := getRatio(earningBeforeTax, earningAfterTax)

	fmt.Printf("Earning before tax = %.2f \n", earningBeforeTax)

	fmt.Printf("Profit = %.2f \n", earningAfterTax)

	fmt.Printf("Ratio = %.2f", ratio)
}

// This format is correct
func getInput(inputData string) float64 {
	var input float64
	prompt := fmt.Sprintf("Enter %s: ", inputData)
	fmt.Printf("%v", prompt)
	fmt.Scanln(&input)
	return input
}

// This format is also correct but not recommended because empty return is not clear and can be confusing for the reader
func getEarning(revenue, expense, taxRate float64) (earningBeforeTax, earningAfterTax float64) {
	earningBeforeTax = revenue - expense
	earningAfterTax = earningBeforeTax - (earningBeforeTax * (taxRate / 100))
	return
}

func getRatio(earningBeforeTax, earningAfterTax float64) float64 {
	ratio := earningBeforeTax / earningAfterTax
	return ratio
}
