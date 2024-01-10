package main

import (
	"fmt"
	"math"
	"strings"
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

func getInitials(n string ) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _,v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) <1 {
		return initials[0], "_"
	}
	return initials[0], initials[1]
}


func main() {
	
	names := []string{"Kenny" , "Punk", "Adam"}
	cycleNames(names,sayGreeting )
	cycleNames(names,sayBye )

	area1 := circleArea(45.5)
	fmt.Printf("CIrcle area : %2.3f \n", area1)

	fn , sn := getInitials("Jeffery Epstien")
	fmt.Printf("Intials : %v %v", fn , sn )


}

