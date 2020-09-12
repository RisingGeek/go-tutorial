package main

import (
	"fmt"
	"math"
	"time"
)

/*
	func functionName(variable name variable type) return type {}
	func functionName(variable name, variable name variable type) return type {} -> if parameters have same value
*/
func add(a int, b int) int {
	return a + b
}

/* Multiple return */
func multiple(a, b string) (string, string) {
	return a, b
}

/* Naked return */
func naked(a int) (nakedVar int) {
	nakedVar = a * 2
	return
}

func main() {
	fmt.Println("Hello, Go")
	fmt.Println("Time", time.Now())
	fmt.Println("Square root of 9 is", math.Sqrt(9))

	fmt.Println("Add function", add(10, 20))
	str1, str2 := multiple("string1", "string2")
	fmt.Println("Return 2 strings", str1, str2)
	fmt.Println("Naked return", naked(10))

	// Initializing variable
	var (
		intVar  int  = 10
		boolVar bool = true
	)
	fmt.Println("Integer variable", intVar, "Boolean variable", boolVar)

	// Short assignment
	shortAssignment := 2
	fmt.Println("Short assignment", shortAssignment)

	/*
		Basic types:
		bool, string, int, int8, int16, int32, int64, float32, float64
	*/

	fmt.Println("Looping starts in another file")
	flowControl()
	types()
}
