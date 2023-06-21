// Reference https://coderwall.com/p/cp5fya/measuring-execution-time-in-go

// Usage example
//
//	func factorial(n *big.Int) (result *big.Int) {
//	    defer timeTrack(time.Now(), "factorial")
//	    // ... do some things, maybe even return under some condition
//	    return n
//	}

package main

import (
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
