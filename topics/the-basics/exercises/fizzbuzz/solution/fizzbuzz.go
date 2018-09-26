// fizzbuzz
// A program that prints the numbers from 1 to n. But for multiples of three
// print “Fizz” instead of the number and for the multiples of five print “Buzz”.
// For numbers which are multiples of both three and five print “FizzBuzz”.
package main

// Add your imports here.
import (
	"fmt"
)

// Could return a string and print it in main
func fizzbuzz(input int) {
	switch {
	case (input%3 == 0 && input%5 == 0):
		fmt.Println("FizzBuzz")
	case (input%3 == 0):
		fmt.Println("Fizz")
	//brackets are optional i just like em
	case input%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(input)
	}
}

func main() {
	//loop through numbers from 1 to 1,000
	for i := 1; i <= 1000; i++ {
		// Use you function to check each value
		fizzbuzz(i)

	}

}
