package main

import (
	"fmt"
)

func main () {
	balance := 1800000
	desiredbalance := 240000

	if balance < desiredbalance {
		fmt.Println("more is coming, just watch God provide")
		if balance < 1000000 {
			fmt.Println("Lokoka")
		}
	} else if balance > desiredbalance {
		fmt.Println("Get born again, He's bottomless")
		if balance < 2000000 {
			fmt.Println("Lokoka")
		}
	} else {
		fmt.Println("Dream bigger borther. this is the winner statement")
	}

	if name, othername := "mine", "alone"; name == othername {
		fmt.Println("yes, that one")
	} else if name != othername {
		fmt.Println("Huh!!")
	} else {
		fmt.Println("Seek his will")
	}
}