package main

func main() {

	listData := []float64{1, 2, 3, 4}

	listData = append(listData, 5, 6, 7)

	newList := []float64{8, 9, 10}

	// Same as spread operator in JavaScript
	listData = append(listData, newList...)
}
