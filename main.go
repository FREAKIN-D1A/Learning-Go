package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "Hello there, my friends"
	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.ReplaceAll(greetings, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.ToLower(greetings))
	fmt.Println(strings.Index(greetings, "th"))

	ages := []int{45,78,98,23,5,47,77,14,10}
	sort.Ints(ages)
	fmt.Println(ages)
}
