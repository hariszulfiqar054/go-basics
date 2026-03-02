package main

import "fmt"

func main() {
	currentBalance := 1000.0
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
		case 2:
			withdraw := 0
			fmt.Print("Withdraw Amount : ")
			fmt.Scanln(&withdraw)
			currentBalance -= float64(withdraw)
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
