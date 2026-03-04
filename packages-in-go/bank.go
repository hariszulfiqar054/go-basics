package main

import (
	"fmt"
	"package-hello.com/filehandling/filehandling"
)

func main() {
	currentBalance, err := filehandling.ReadFromFile()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
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
			filehandling.WriteToFile(currentBalance)
		case 2:
			withdraw := 0
			fmt.Print("Withdraw Amount : ")
			fmt.Scanln(&withdraw)
			currentBalance -= float64(withdraw)
			filehandling.WriteToFile(currentBalance)
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
