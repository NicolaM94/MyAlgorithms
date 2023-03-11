package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	arraymanipulation "myalgo/arrayManipulation"
	"os"
	"strconv"
)

func main() {

	buffs, err := os.ReadFile("/home/nicola/Desktop/grid.csv")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(bytes.NewReader(buffs))
	var grid [][]int
	var line []int

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		for entry := range record {
			converted, _ := strconv.Atoi(record[entry])
			line = append(line, converted)
			if entry%20 == 0 {
				grid = append(grid, line)
				line = []int{}
			}
		}
	}

	fmt.Println(grid)

	test := []int{0, 1, 2, 3, 4, 5}
	test = arraymanipulation.LeftShift(test)
	fmt.Println(test)
	test = arraymanipulation.RightShift(test)
	fmt.Println(test)
	test = arraymanipulation.DeleteIndex(1, test)
	fmt.Println(test)

	testArrayString := []string{"a", "b", "c", "d"}
	testArrayInt := []int{0, 1, 2, 3, 4, 5, 6, 7}

	testArrayString = arraymanipulation.DeleteElem(testArrayString, "a")
	fmt.Println(testArrayString)

	testArrayInt = arraymanipulation.DeleteElem(testArrayInt, 4)
	fmt.Println(testArrayInt)
}
