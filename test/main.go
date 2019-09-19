package main

import "go_set"

func main() {
	s1 := Set{elems: []int{1, 1, 2, 3, 4, 2, 1}}
	fmt.Println(s1.Elements())
}
