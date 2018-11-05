package main
import (
	"fmt"
	"strings"
)

func main() {
	first := "scott"
	last := "businge"

	fmt.Println(convert(first, last))
}

func convert (first, last string) (s1, s2 string) {
	first = strings.Title(first)
	last = strings.ToUpper(last)

	return first, last
}