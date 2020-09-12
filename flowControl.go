package main

import "fmt"

func flowControl() {
	looping()
	conditionals()
	deferExample()
	deferStack()
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

	// Switch statement
	variable := 2
	switch variable {
	case 1:
		fmt.Println("switch 1")
	case 2:
		fmt.Println("switch 2")
		fmt.Println("print line 2")
	default:
		fmt.Println("switch none")
	}
}

func deferExample() {
	// Defer statement does not execute function until the surrounding function returns
	defer fmt.Println("Defered World")
	fmt.Println("Hello")
}

func deferStack() {
	// Defered stack
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
