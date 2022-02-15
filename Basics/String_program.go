package main

import "fmt"

func Concat(data string, size int) (output string) {
	return data[:size]
}

func main() {
	//a := 5
	b := "Welcome"
	c := Concat(b, 3)
	fmt.Println("Hello world, this is my text : ", c)
}
