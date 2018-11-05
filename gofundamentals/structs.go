package main

import (
	"fmt"
)

func main() {
	type Bucketlist struct {
		Name string
		Item string
		Rate int
	} 
// declaration
// 	var Travel Bucketlist
// declaration
// 	Travel := new(Bucketlist)

	Travel := Bucketlist {
		Name: "Partner",
		Item: "visit denver",
		Rate: 7,
	}

	fmt.Println(Travel)
}