package main

import (
	"fmt"
	"reflect"
	"os"
)

var (
	name string
	course string
	module float64
	age, downloads = 4, 600
)

func main () {

	name := os.Getenv("USER")
	series := "designated survivor"

	const status = "born again"

	a := 10.5
	b := 5
	ptr := &b

	fmt.Println("name is ", name, "and of type", reflect.TypeOf(name))
	fmt.Println("course is ", course, "and of type", reflect.TypeOf(course))
	fmt.Println("module is ", module, "and of type", reflect.TypeOf(module))
	fmt.Println("age is ", age, "and of type", reflect.TypeOf(age))
	fmt.Println("memory address of name is", ptr, "and it's value is", *ptr)
	c := int(a) + b

	fmt.Println("C is of value:", c, "and is of type", reflect.TypeOf(c))
	fmt.Println("Hi", name, "you are watching", series)
	changeMovie(&series)
	fmt.Println("Hi", name, "you are now watching", series)

	fmt.Println("Hi", name, "you are forever", status)

	for _, env := range os.Environ() {
	fmt.Println(env)
	}
}

func changeMovie (series *string) string{
	*series = "dynasty"
	fmt.Println("you are watching", *series)
	return *series
}
