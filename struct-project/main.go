package main

import (
	"fmt"

	"haris-blog.com/blog/blog"
)

func main() {
	blogTitle := getInput("Input Blog Title")
	blogContent := getInput("Input Blog Content")

	blogData, err := blog.New(blogTitle, blogContent)
	if err != nil {
		fmt.Println("Error creating blog:", err)
		return
	}

	blogData.DisplayContent()
}
func getInput(inputType string) string {
	var value string
	fmt.Printf("%v : ", inputType)
	fmt.Scanln(&value)
	return value
}
