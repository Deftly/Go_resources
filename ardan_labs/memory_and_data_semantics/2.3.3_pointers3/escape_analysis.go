package main

import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
}

func main() {
  u1 := createUserV1()
  u2 := createUserV2()

  println("u1", &u1, "u2", &u2)
}

// createUserV1 creates a user value and passes
// a copy back to the caller.
func createUserV1() user {
  u := user {
    name: "Bill",
    email: "bill@ardanlabs.com",
  }

  fmt.Println("V1", &u)
  return u
}

func createUserV2() *user {
  u := user{
    name: "Bill",
    email: "bill@ardanlabs.com",
  }

  fmt.Println("V2", &u)
  return &u
}
