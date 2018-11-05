// stand alone exec program rather than shared library
package main

import (
	"fmt"
)

// entry program within package main
func main () {
	Cars := []string {"Mercedes", "Audi", "Ferrari"}
	Caz := []string {"BMW", "Tesla", "Ford", "Mercedes"}

	for _, i := range Cars{
		for _, j := range Caz{
			if i == j {
				fmt.Println("Similar cars found.", i)
			}
		}
	}
}