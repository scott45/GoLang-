package main

import (
	"fmt"
)

func main () {
	Networth := allWealth(30, 23, 67, 89, 345, 567)
	fmt.Println(Networth)
}

func allWealth (monies ...int) int{
	most := monies[0]
	for _, i := range monies {
		if i > most {
			most = i
		}
	}
	return most
}