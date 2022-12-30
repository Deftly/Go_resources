package main

import "fmt"

func main() {
  // Declare an array of 5 strings initialized with values
  friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

  // Iterate over the array displaying the value and address of each element
  for i, v := range friends {
    fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &friends[i])
  }

  /* Using the value semantic form of the for range.
  The reason Betty is printed instead of Jack is because this for range
  isn't iterating over the friends array but rather its own copy, so any
  changes to friends is irrelevant in this iteration */
  fmt.Printf("Bfr[%s] : ", friends[1])
  for i, v := range friends {
    friends[1] = "Jack"

    if i == 1 {
      fmt.Printf("v[%s]\n", v)
    }
  }

  fmt.Printf("v[%s]\n", friends[1]) // Jack
}
