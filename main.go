package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("numbers")
	if err != nil {
		panic(err)
	}
	var reader []string
	var tempString string
	for c := range file {
		if string(file[c]) != "\n" {
			tempString += string(file[c])
		}
		if len(tempString) == 50 {
			reader = append(reader, tempString)
			tempString = ""
		}
	}
	for v:=range reader {
		fmt.Println(reader[v])
	}


	var outDigits [10]int
	counter := len(reader[0])-1

	


}
