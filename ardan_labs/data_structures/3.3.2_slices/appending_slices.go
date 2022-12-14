package main

import "fmt"

func main() {
  // Declare a nil slice of strings. This means there is no backing array
  var data []string // The slice pointer points to the empty struct: var es struct{}
  // data := make([]string, 0, 1e5), because we know how much capacity we need we initialize the backing array at the start

  // Get the capacity of the slice.
  lastCap := cap(data)

  // Append ~100k strings to the slice.
  for record := 1; record <= 1e5; record++ {
    value := fmt.Sprintf("Rec: %d", record)
    data = append(data, value)

    // When the capacity of the slice changes, display the changes.
    if lastCap != cap(data) {
      // Calculate the percent of change.
      capChg := float64(cap(data)- lastCap) / float64(lastCap) * 100

      lastCap = cap(data)

      // Display the results.
      fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
              &data[0],
              record,
              cap(data),
              capChg)
    }
  }
}
