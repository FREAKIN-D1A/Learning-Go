package main

import ( 
"fmt"
 "strings"
)

func main() {
	// arrays
	var ages [3]int = [3]int{20, 21, 22}
	ages[2] = 56
	fmt.Println(ages)

	var ages2 = [3]int{20, 21, 22}
	fmt.Println(ages2)

	names := [4]string{"Yosi", "Tanahasi", "Omega", "ALpha"}
	fmt.Println(names, len(names))

	// slices
	var scores = []int{20, 21, 22}
	scores[2] = 25
	fmt.Println(scores, len(scores))

	scores = append(scores,99,78, 45, 78 ,78)
	fmt.Println(scores, len(scores))

	rangeOne := scores[0:3]
	fmt.Println(rangeOne, len(rangeOne))

	
	rangeTwo := scores[2:]
	fmt.Println(rangeTwo, len(rangeTwo))
}
