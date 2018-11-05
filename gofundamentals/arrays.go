//  array is numbered list of same type (strings or ints). fixed length.
// slices more used than arrays. numbered list of a single type. can be resied.
// slices are built on top of an array.
// slices are references. 

package main

import (
	"fmt"
)

func main () {
	Word := []string {"John", "Luke", "Romans", "Timothy"}

	fmt.Printf("Length is: %d. \nCapacity is: %d", 
		len(Word), cap(Word))

	Numbers := []int {0, 1, 2, 3, 4 ,5, 6, 7, 8, 9, 10}
	fmt.Println(Numbers)

	Numbers[3] = 7
	fmt.Println(Numbers)

	newNumbers := Numbers[4:8]
	fmt.Println(newNumbers)
}