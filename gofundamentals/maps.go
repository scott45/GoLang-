package main

import (
	"fmt"
)

func main() {
	// declaration
	wealth := make(map[string]int)
	wealth["scott"] = 564

	marrieds := map[string]string {
		"scott": "crush",
		"businge": "her again",
	}

	fmt.Println("marrieds are: ", marrieds)

}