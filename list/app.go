package main

import "fmt"

/*
	Slice doesn't create a new copy of the array, it just creates a new view of the same array.
	So when you modify the slice, it modifies the underlying array. This is why the capacity of the slice is less than the capacity of the original array,
	because it starts from the index where the slice begins.
*/

func main() {
	prices := [5]float64{1, 2, 3, 4, 5}
	featuredPrices := prices[:4]
	moreData := featuredPrices[2:]
	fmt.Println(featuredPrices,cap(moreData))
}
