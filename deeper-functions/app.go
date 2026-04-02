/*
	- Function can be used as a value same as js/ts.

*/

package main

import "fmt"

type FunctionType func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	transformerFun := getTransformerFunc("double")
	transformedData := transformNumber(&numbers, transformerFun)
	// This is the anonymous function, it is a function without a name.
	transformedData2 := transformNumber(&numbers, func(value int) int {
		return value * 3
	})
	fmt.Print(transformedData)
	fmt.Print(transformedData2)

}

func transformNumber(numbers *[]int, transformer FunctionType) []int {
	transformedList := []int{}
	for _, value := range *numbers {
		transformedList = append(transformedList, transformer(value))
	}
	return transformedList
}

func double(value int) int {
	return value * 2
}

func getTransformerFunc(funcType string) FunctionType {
	switch funcType {
	case "double":
		return double
	default:
		return nil
	}
}
