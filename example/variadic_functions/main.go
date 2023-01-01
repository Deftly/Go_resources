package main

import "fmt"

// Here's a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
  fmt.Print(nums, " ")
  total := 0

  // Within the function, the type nums is equivalent to []int.
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func main() {
  // Variadic functions can be called in the usual way with individual arguments.
  sum(1, 2)
  sum(1, 2, 3)

  // If you have multiple args in a slice, apply them to a variadic function
  // using func(slice...) like below.
  nums := []int{1, 2, 3, 4}
  sum(nums...)
}
