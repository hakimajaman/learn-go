// arrays
// 	creation
// 	built-in functions
// 	working with arrays

// slices
// 	creation
// 	built-in functions
// 	working with slices

package main

import "fmt"

func main() {
	// we need to set the size of the array first,
	// and array only can input 1 type.
	grades := [3]int{12, 23, 34}

	fmt.Println(grades)

	// if you want to have more than 1 type in array,
	// you can set it to string type.

	// if you don't know the size of the array,
	// you can input 3 dots.
	brothers := [...]int{32,322,233,4312}
	fmt.Println(brothers)

	// in array, index always start from 0
	var students [4]string // here you only set the variable
	fmt.Println(students) // you will get [  ], let say, the data is only spacebar

	// if you want to push or add the data,
	// you can input it by index
	students[0] = "Rifqi" // we can say it's edit the index 0
	fmt.Println(students)

	// to check length of the array
	fmt.Println(len(students))

	// to set array in array
	var teachers [3][3]string
	teachers[0] = [3]string{"English","Bahasa Indonesia","Bahasa Melayu"}
	fmt.Println(teachers)


	fmt.Println("----------------")
	// lets start slicing
	slicing := [...]int{1,2,3,4,5,6,7,8,9}
	a := slicing[:] // slice all elements
	b := slicing[3:] // slice from 4th element to end
	c := slicing[:6] // slice from first 6 elements
	d := slicing[3:6] // slice 4th, 5th, and 6th elements

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)


	// you only can input the data to the capacity maximum size
	withCapacity := make([]int, 0, 100)
	fmt.Printf("it's value %v, and it's capacity, %d\n", withCapacity, cap(withCapacity))

}
