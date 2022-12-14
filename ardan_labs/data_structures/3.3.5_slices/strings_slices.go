package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
  // Declare a string with both chinese and english characters.
  s := "世界 means world"

  // UTFMax is 4 -- up to 4 bytes per encoded rune.
  var buf [utf8.UTFMax]byte

  // Iterate over the string.
  for i, r := range s {
    // Capture the number of bytes for this rune.
    rl := utf8.RuneLen(r)

    // Calculate the slice offset for the bytes associate with this rune.
    si := i + rl
    
    // Copy of the rune from the string to our buffer.
    copy(buf[:], s[i:si])

    // Display the details.
    fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
  }
}
