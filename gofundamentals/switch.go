// stand alone exec program rather than shared library
package main

import (
	"fmt"
)

// entry program within package main
func main () {
	switch "scott" {
	case "businge":
		fmt.Println("hello scott")
	case "scott":
		fmt.Println("hello managah")
	case "abraham":
		fmt.Println("hello scott again")
	case "amooti":
		fmt.Println("hello scott again again")
	default:
		fmt.Println("hello mheenn")
	}
}