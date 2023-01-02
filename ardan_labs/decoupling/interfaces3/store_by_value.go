package main

import "fmt"

// printer displays information.
type printer interface {
  print()
}

// cannon defines a cannon printer.
type cannon struct {
  name string
}

// print displays the printer's name.
func (c cannon) print() {
  fmt.Printf("Printer Name: %s\n", c.name)
}

// epson defines an epson printer.
type epson struct {
  name string
}

func (e *epson) print() {
  fmt.Printf("Printer Name: %s\n", e.name)
}

func main() {
  // Create cannon and epson printer.
  c := cannon{"PIXMA TR4520"}
  e := epson{"WorkForce Pro WF-3720"}

  // Add the printers to the collection using both value
  // and pointer semantics.
  printers := []printer{
    c,
    &e,
  }

  // Change the name field for both printers.
  c.name = "PROGRAF PRO-1000"
	e.name = "Home XP-4100"

  // Iterate over the slice of printers and call print against the
  // cpied interface value.
  for _, p := range printers {
    p.print()
  }

  // When we store a value, the interface value has its own copy of the value.
  // Change to the original value will will not be seen.

  // Whenw e store a pointer, the itnerface value has its own copy of the address.
  // Changes to the orginal value will be seen. 
}
