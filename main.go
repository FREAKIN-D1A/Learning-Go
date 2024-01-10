package main

import (
	"fmt"
	"math"
)


func sayGreeting(name string  ){
	fmt.Printf("Good Morning , %v \n", name)
}

func sayBye(name string  ){
	fmt.Printf("Good bye , %v \n", name)
}

func cycleNames( names []string  , fun func(string)  ) {
	for _, name := range names {
		fun(name)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}


func main() {
	
	names := []string{"Kenny" , "Punk", "Adam"}
	cycleNames(names,sayGreeting )
	cycleNames(names,sayBye )

	area1 := circleArea(45.5)
	
	fmt.Printf("CIrcle area : %2.3f \n", area1)
}

