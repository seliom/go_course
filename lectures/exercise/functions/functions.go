//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greetName(name string) {
	fmt.Println("Hey " + name + " :)")
}
//* Write a function that returns any message, and call it from within
//  fmt.Println()
func outputRandomMessage() string{
	return "Random message"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func tripleSum(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

//* Write a function that returns any number
func outputRandomNumber() int {
	return 7
}

//* Write a function that returns any two numbers
func outputTwoRandomNumbers() (int, int) {
	return 7,1
}

//* Add three numbers together using any combination of the existing functions.
//  * Print the result

//* Call every function at least once

func main() {
	greetName("Max")
	fmt.Println(outputRandomMessage())
	num1, num2 := outputTwoRandomNumbers()
	sum := tripleSum(num1, num2, outputRandomNumber())
	fmt.Println(sum)

}
