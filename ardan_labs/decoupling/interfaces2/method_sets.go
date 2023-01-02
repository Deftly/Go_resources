package main

import "fmt"

// notifier is an interface that defines notification type behavior.
type notifier interface {
  notify()
}

type user struct {
  name string
  email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
  fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func main() {
  // Create a value of type User and send a notification.
  u := user{name: "Bill", email: "bill@email.com"}
  fmt.Println(u)
  // Values of type user do not implement the interface because
  // pointer receivers don't belong to the method set of a value.

  sendNotification(&u)
  /* cannot use u (type user) as type notifier in argument to sendNotification:
  user does not implement notifier (notify method has pointer receiver) */
}

func sendNotification(n notifier) {
  n.notify()
}
