package main

import (
	"fmt"
	arraymanipulation "myalgo/arrayManipulation"
	"os"
	"strconv"
)

func main() {

	// Grabbing numbers from text file and storing them in matrix
	file, err := os.ReadFile("numbers")
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

	outArray := []int{0,0,0,0,0,0,0,0,0,0}
	var remainder int
	var counter int = len(matrix[0])-1

	for counter >= 0 {
		result := 0 
		for value := range matrix {
			result += int(matrix[value][counter] - 48)
		}
		result += remainder
		stringResult := fmt.Sprint(result)
		outArray = arraymanipulation.ShiftRightToInsert(outArray[:],int(stringResult[len(stringResult)-1]-48))
		remainder,_ = strconv.Atoi(string(stringResult[:len(stringResult)-1]))
		counter--
	}
	counter = len(fmt.Sprint(remainder))-1
	for counter >= 0 {
		outArray = arraymanipulation.ShiftRightToInsert(outArray,int(fmt.Sprint(remainder)[counter]-48))
		counter--
	}
	
	fmt.Println(outArray)

	fmt.Println(calculateHugeSums("numbers",10))

}

// Used to perform the sum of the numbers in filePath up to firstNdigits precision.
// The numbers to sum must be held by a file, one number for each row
// It needs the function ShiftRightToInsert, which you can find it here
func calculateHugeSums (filePath string, firstNdigits int) []int {

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

	//Populating outArray with the required nth precision
	outArray := []int{}
	for n:=0;n < firstNdigits;n++ {
		outArray = append(outArray, 0)
	}

	// Remainder variable used to store the remainder of each column addition
	var remainder int

	// Counter keeps track of the current column of the matrix, starting from the end
	var counter int = len(matrix[0])-1

	// Loop over the columns, starting from the end
	for counter >= 0 {
		// Init a result variable equal to 0 to store the column result
		result := 0 
		// Add each value of the column[counter] to the result
		// NOTE: Add matrix[value][counter] - 48 because matrix[value] is a rune, where 0 = 48
		for value := range matrix {
			result += int(matrix[value][counter] - 48)
		}
		// Add the remainder of the previous column addition
		result += remainder
		// Turns the result into a string, to grab the last digit and the next remainder
		stringResult := fmt.Sprint(result)
		// Add the last digit of the result to the outArray
		outArray = arraymanipulation.ShiftRightToInsert(outArray[:],int(stringResult[len(stringResult)-1]-48))
		// Rewrite remainder as the result with the last digit left over
		remainder,_ = strconv.Atoi(string(stringResult[:len(stringResult)-1]))
		// Goes to the previous column
		counter--
	}
	// When over, the process still needs to add the final remainder in front of the array
	counter = len(fmt.Sprint(remainder))-1
	for counter >= 0 {
		outArray = arraymanipulation.ShiftRightToInsert(outArray,int(fmt.Sprint(remainder)[counter]-48))
		counter--
	}
	
	return outArray
}