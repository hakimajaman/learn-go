// boolean type
// numeric types
// 	integer
// 	floating point
// 	complex number
// text types
package main

import (
	"fmt"
)

func main() {
	var n bool = true
	fmt.Printf("its %v, type is %T\n", n, n)

	// compare system
	j := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", m, m)
}
