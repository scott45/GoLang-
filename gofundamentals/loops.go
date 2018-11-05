// stand alone exec program rather than shared library
package main

import (
	"fmt"
	"time"
)

// entry program within package main
func main () {
	for timer := 20; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Bboooomm!!!!")
			break
		}
	fmt.Println(timer)
	time.Sleep(1 * time.Second)
	}
}