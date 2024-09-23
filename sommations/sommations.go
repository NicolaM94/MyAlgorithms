package sommations

import (
	"fmt"
	arraymanipulation "myalgo/arrayManipulation"
	"os"
	"strconv"
)

// Simple sum of an array of integers
func IntArraySum(array []int) (sum int) {
	for v := range array {
		sum += array[v]
	}
	return
}

// Used to perform the sum of the numbers in filePath up to firstNdigits precision.
// The numbers to sum must be held by a file, one number for each row
// It needs the function [ShiftRightToInsert], which you can find it here
func CalculateHugeSums(filePath string, firstNdigits int) []int {
	// Grabbing numbers from text file and storing them in matrix
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var matrix []string
	var tempString string
	for c := range file {
		if string(file[c]) != "\n" {
			tempString += string(file[c])
		}
		if len(tempString) == 50 {
			matrix = append(matrix, tempString)
			tempString = ""
		}
	}
	// Real implementation
	outArray := []int{}
	for n := 0; n < firstNdigits; n++ {
		outArray = append(outArray, 0)
	}
	var remainder int
	var counter int = len(matrix[0]) - 1
	for counter >= 0 {
		result := 0
		for value := range matrix {
			result += int(matrix[value][counter] - 48)
		}
		result += remainder
		stringResult := fmt.Sprint(result)
		outArray = arraymanipulation.ShiftRightInsertLeft(outArray[:], int(stringResult[len(stringResult)-1]-48))
		remainder, _ = strconv.Atoi(string(stringResult[:len(stringResult)-1]))
		counter--
	}
	counter = len(fmt.Sprint(remainder)) - 1
	for counter >= 0 {
		outArray = arraymanipulation.ShiftRightInsertLeft(outArray, int(fmt.Sprint(remainder)[counter]-48))
		counter--
	}
	return outArray
}
