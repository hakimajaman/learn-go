// naming convention
// typed constants
// enumerated contants
// enumerations expressions

package main

import "fmt"

const toShadow = 20

// we can create an index number by using iota
const (
	a = iota
	b = iota
	c = iota
)

// the e and f will automatically use same as before,
// it's d, so e and f will be iota too.
const (
	d = iota
	e
	f
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	const myConst int = 42
	// myConst = 5 // we cannot change the constant
	// but we still can shadow it
	const toShadow int = 32

	fmt.Printf("here's constant %v with type %T\n", myConst, myConst)
	fmt.Printf("here's the shadow %v with type %T\n", toShadow, toShadow)

	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 2

	fmt.Println(d) // 0
	fmt.Println(e) // 1
	fmt.Println(f) // 2

	fmt.Println("------------------")

	const roles byte = isAdmin | canSeeAfrica | canSeeAsia
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin)
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == canSeeAsia)
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == canSeeEurope)
}


// SUMMARY
// immutable, but can be shadowed

// replaced by the compiler at compile time
// 	value must be calculable at compile time

// typed constants work like immutable variables
// 	can interoperate only with same type

// untyped constants work like literals
// 	can interoperate with similar type

// enumerated constants
// 	special symbol iota allowed related constants to be created easily
// 	iota start at 0 in each const block and increments by one
// 	watch out of constant values that match zero values for variables

// enumerated expressions
// 	operations that can be determined at compile time are allowed
// 		arithmetic
// 		bitwise operation
// 		bitshifting
