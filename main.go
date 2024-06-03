package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, Wolrd!")

	var intNum int = 10
	fmt.Println(intNum)

	var floatNum float32 = 3.1415
	fmt.Println(floatNum)

	var myString string = "Hello\nWorld"
	fmt.Println(myString)

	myVar := 10
	fmt.Println(myVar)

	printMe("I am izac")

	var result, remainder, err = intDivision(10, 0)
	if err != nil {
		fmt.Println(err)
	} else if remainder == 0 {
		fmt.Printf("The result of integer division is %v", result)
	} else {
		fmt.Printf("The result of integer division is %v with remainder %v", result, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, nil
}
