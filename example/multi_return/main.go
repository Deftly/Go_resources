package main

import "fmt"

// Here the function signature tells us that the function returns two ints
func vals() (int, int) {
  return 3, 7
}

func main() {
  a, b := vals()
  fmt.Println(a)
  fmt.Println(b)

  _, c := vals()
  fmt.Println(c)
}
