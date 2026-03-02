package main

import (
	"fmt"
	"os"
	"strconv"
)

func writeToFile(balance float64) {
	balanceData := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceData), 0644)
}

func readToFile() float64 {
	fileData, err := os.ReadFile("balance.txt")
	if err == nil {
		parsedString := string(fileData)
		data, _ := strconv.ParseFloat(parsedString, 64)
		return data
	}
	return 0
}

func main() {
	currentBalance := readToFile()
	for {

		fmt.Println(`
Welcome to the Bank!
1. Deposit
2. Withdraw
3. Check Balance
4. Exit
`)
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			depositData := 0
			fmt.Print("Deposit Amount : ")
			fmt.Scanln(&depositData)
			currentBalance += float64(depositData)
			writeToFile(currentBalance)
		case 2:
			withdraw := 0
			fmt.Print("Withdraw Amount : ")
			fmt.Scanln(&withdraw)
			currentBalance -= float64(withdraw)
			writeToFile(currentBalance)	
		case 3:
			fmt.Printf("Your balance : %v\n", currentBalance)
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
