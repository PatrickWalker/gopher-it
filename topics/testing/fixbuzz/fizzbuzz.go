// fizzbuzz
// A program that prints the numbers from 1 to n. But for multiples of three
// print “Fizz” instead of the number and for the multiples of five print “Buzz”.
// For numbers which are multiples of both three and five print “FizzBuzz”.
package main

// Add your imports here.
import (
	"fmt"
	"strconv"
)

// Could return a string and print it in main
func fizzbuzz(input int) (val string) {
	if input%3 == 0 || input%5 == 0 {
		val = "FizzBuzz"
	} else if input%3 == 0 {
		val = "Fizz"
	} else if input%5 == 0 {
		val = "Buzz"
	} else {
		val = strconv.Itoa(input)
	}
	return
}

func main() {
	//loop through numbers from 1 to 1,000
	for i := 1; i <= 30; i++ {
		// Use you function to check each value
		fmt.Println(fizzbuzz(i))

	}

}
