/*
	Interface in go allow us to define a set of methods that a type must implement.
	It is a powerful tool for achieving polymorphism and abstraction in Go.
	An interface defines a contract that a type must fulfill, without specifying how the type should implement the methods.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = saveData(userNote)

	

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func saveData(s saver) error {
	err := s.Save()

	if err != nil {
		return err
	}

	return nil
}
