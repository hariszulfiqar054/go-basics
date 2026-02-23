/*
In the go app, go program much organized into packages, you can have multiple packages

	in a single program and can have different packages organized in different files.

	The main is special package name that this package is the main entry point of the program.

	The function should also be named as main because go will call and execute that function when the program is run.

	Only one main function is allowed in the main package, and that is the entry point of the program.

	Its important to create the build of go app and make it executable, to do that we need to have main package and main function in it.

	The go build command need to be run the directory where the module exists, to create module you need to define module path

		go mod init <module-name>

	after that you can run the go build command to create executable file of the app.
*/
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
