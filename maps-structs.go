// maps
// creating
// manipulation

// structs
// creating
// naming conventions
// embedding
// tags

package main

import "fmt"

type Doctor struct {
	number int
	name string
}

func main() {
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas": 27862596,
		"Florida": 20612439,
		"Pennsylvania": 19745289,
		"Illinois": 12802503,
		"Ohio": 11614373,
	}
	fmt.Println(statePopulations);

	// if you don't have a data and want to create a map
	// you can set the variable first using make
	myMap := make(map[string]int)

	// and then you can add it later;
	// usually it use for looping and push.
	myMap = map[string]int{
		"john": 1212,
	}

	statePopulations["Ohio"] = 1929
	delete(statePopulations, "Pennsylvania")
	fmt.Println(statePopulations);

	fmt.Println(myMap)
	fmt.Println(myMap["john"])

	thisDoctor := Doctor {
		number: 3,
		name: "Hakim",
	}

	fmt.Println(thisDoctor)
}
