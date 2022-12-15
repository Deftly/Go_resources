package main

import "fmt"

func main() {
	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")

	// Pass the "value of" the count.
	increment(count)

	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")
}

// increment takes in a variable and increments it by 1.
func increment(inc int) {
  // Increment the "value of" inc.
  inc++
  fmt.Println("inc:\tValue of[", inc, "]\tAddr of[", &inc, "]")
}
