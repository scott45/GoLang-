// stand alone exec program rather than shared library
package main

import (
	"fmt"
	"os"
)

// entry program within package main
func main () {
	_, err := os.Open("c:\\temp\\file.txt")

	if err != nil {
		fmt.Println("error returned was:", err )
	}
}