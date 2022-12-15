package main


import "fmt"

func main() {
	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	fmt.Println("count:\tValue of[", count, "]\tAddr of[", &count, "]")

	// Pass the "address of" the count.
	increment(&count)

	fmt.Println("count:\tValue of[", count, "]\tAddr of[", &count, "]")
}

func increment(inc *int) {
  // Increment the "value of" count that the "pointer points to".
  *inc++

  fmt.Println("ince:\tValue of[", inc, "]\tAddr of[", &inc, "]\tValue Points To[", *inc, "]")
}
