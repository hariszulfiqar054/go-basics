package main

import "fmt"

func main() {
	hobbies := []string{"Eating", "Sleeping", "Gym"}
	fmt.Println(hobbies)
	fmt.Printf("First element %v\n", hobbies[0])
	fmt.Printf("Other elements %v\n", hobbies[1:])

	newSlice := hobbies[:2]
	fmt.Println(newSlice)

	courseGoal := []string{"Travel", "Food"}
	courseGoal[1] = "Learn Go"
	fmt.Println(courseGoal)
	courseGoal = append(courseGoal, "Learn Python")
	fmt.Println(courseGoal)

}
