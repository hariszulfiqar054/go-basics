/*
 In order to make the split code into a different module you need a new folder
 with the name of that package.

 In order to export the functions or variable the name should start with a
 capital letter. This is convention in Go that must be followed

 In order to import and use that module you need to import with the whole module path
 like this : modulename/filename
*/

package filehandling

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteToFile(balance float64) {
	balanceData := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceData), 0644)
}

func ReadFromFile() (float64, error) {
	fileData, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0, errors.New("File doesn't exist")
	}
	parsedString := string(fileData)
	data, _ := strconv.ParseFloat(parsedString, 64)
	return data, nil
}
