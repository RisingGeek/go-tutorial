package main

import "fmt"

func flowControl() {
	looping()
	conditionals()
}

func looping() {
	// Go has only for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func conditionals() {
	a := true
	if a {
		fmt.Println("Statement true")
	}
}
