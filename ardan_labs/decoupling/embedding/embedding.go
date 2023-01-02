package main

import "fmt"

// notifier is an interface that defines notification type behavior.
type notifier interface {
  notify()
}

// user defines a user in the program
type user struct {
  name string
  email string
}

// notify is a method that notifies users of different events.
func (u *user) notify() {
  fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// admin represents an admin user with privileges.
type admin struct {
  user // Embedded type
  level string
}

func main() {
  // Create an admin user.
  ad := admin{
    user: user{
      name: "john smith",
      email: "john@email.com",
    },
    level: "super",
  }

  // We can access the inner type's method directly.
  // ad.user.notify()
  // The inner type's method is promoted.
  // ad.notify()

  sendNotification(&ad)
}

func sendNotification(n notifier) {
  n.notify()
}
