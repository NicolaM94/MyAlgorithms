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

	container := []string{}
	temp := ""

	for _, v := range file {
		if v == 10 {
			container = append(container, temp)
			temp = ""
		}
		temp += string(v)
	}
	container = append(container, temp)

	fmt.Println(len(container))
}
