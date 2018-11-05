// stand alone exec program rather than shared library
package main

import (
	"fmt"
	"runtime"
)

// entry program within package main
func main () {
	fmt.Println("hello world", runtime.GOOS)
}