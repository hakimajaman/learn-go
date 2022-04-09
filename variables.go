// variable declaration
// redeclaration and shadowing
// visibility
// naming conventions
// type conventions

package main

import (
	"fmt"
	"strconv" // you need to add this to convert number to strings
)

var (
	fullName string = "Rifqi Taufiqul Hakim"
	gender string = "Male"
)

// we can use 3 types for the variable struct
var i int // if you don't want to input something first, and you pretty sure it's integer
var j int = 30 // you can input the number directly

// you can export your variable outside the function,
// by naming the variable first letter with the uppercase
var I int = 40

// if you declare the variable inside the function,
// you need to use it or error.
func main() {
	// you only can use this struct variable inside function,
	// and you can use it if you doesn't know what's type of it.
	k := "it's string"
	fmt.Println(k)


	// here's we redeclare the variable from top
	fmt.Println(j) // it will call 30
	var j = 23
	fmt.Println(j) // it will call 23
	// but redeclare cannot use :=

	// to convert the type
	var thisINT int = 23
	fmt.Printf("%v, %T\n", thisINT, thisINT)

	var thisFloat float32 = float32(thisINT)
	fmt.Printf("%v, %T\n", thisFloat, thisFloat)

	var thisString string = strconv.Itoa(thisINT)
	fmt.Printf("%v, %T\n", thisString, thisString)
}

// SUMMARY
// variable declaration
// 	var foo int
// 	var foo int = 42
// 	foo := 42

// can't redeclare variables, but can shadow them

// all variables must be used

// visibility
// 	lowercase first letter for package scope
// 	uppercase first letter to export
// 	no private scope

// naming conventions
// 	Pascal or camelCase
// 		capitalize acronyms(HTTP, URL)
	
// 	as short as reasonable
// 		longer name for longer lives
