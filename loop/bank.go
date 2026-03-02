package main

import "fmt"

/*
- Only for loop is supported in Go.
- Infinite loop can be created using for loop without any condition i.e for {}
- Loop with the condition can be created using for loop with condition i.e for condition {}
- Loop with traditional syntax of for loop can also be created like for i:=0 ; i<10 ; i++ {}
*/
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
		if choice == 1 {
			depositData := 0
			fmt.Print("Deposit Amount : ")
			fmt.Scanln(&depositData)
			currentBalance += float64(depositData)
		} else if choice == 2 {
			withdraw := 0
			fmt.Print("Withdraw Amount : ")
			fmt.Scanln(&withdraw)
			currentBalance -= float64(withdraw)

		} else if choice == 3 {

			fmt.Printf("Your balance : %v\n", currentBalance)
		} else if choice == 4 {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
