package main

import "fmt"

func types() {
	pointers()
	structure()
	arrays()
	rangeArray()
	maps()
}

func pointers() {
	var p *int
	i := 10
	p = &i
	fmt.Println("Pointer address", p, "Pointer value", *p)
	// Go doesn't have pointer arithmetic
}

// Student : A comment should be there with the name of thing that's being exported. Otherwise, linter will give error
type Student struct {
	name string
	roll int
}

func structure() {
	var stud Student = Student{"Bhavay", 1}
	fmt.Println("structure", stud)

	// Pointers to structure
	p := &stud
	fmt.Println("pointer to structure", p.name)

	var stud2 = Student{name: "Bhavay"}
	fmt.Println("Structure without 1 value", stud2)
}

func arrays() {
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)

	arr2 := [2]string{"first", "second"}
	fmt.Println(arr2)

	orig := [5]int{1, 2, 3, 4, 5}
	sliced := orig[1:4] // OR var slice []int = orig[1:4]
	fmt.Println("Sliced array", sliced)
	/*
		Default low is 0 -> [:3]
		Default high is length of array -> [3:]
	*/

	fmt.Println("Appending to slice")
	var s []int
	s = append(s, 1)
	fmt.Println(s)
	s = append(s, 2, 3, 4)
	fmt.Println(s)

}

func rangeArray() {
	fmt.Println("Range loop:")
	var arr = []int{1, 2, 3, 4, 5}
	for i, val := range arr {
		fmt.Println(i, val)
	}

	/*
		Omit variables
		for i,_ := range arr
		for _,val := range arr
		for i := range arr
	*/
}

func maps() {
	var hashMap map[string]int
	hashMap = make(map[string]int)
	fmt.Println("initial map", hashMap)
	hashMap["key1"] = 10
	hashMap["key2"] = 20
	fmt.Println(hashMap)

	/*
		Delete from map
		delete(mapName, key)
	*/

	/*
		Check whether key is present in map or not
		elem, ok := mapName[key]
		ok is true if key is present. elem -> value
		if ok is false, elem is 0 value of its type
	*/
}
