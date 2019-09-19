package go_set

import "fmt"

// Set returns set from list
type Set struct {
	elems []int
}

func (s *Set) Elements() []int {
	tmpMap := make(map[int]bool)
	outList := []int{}
	for _, v := range s.elems {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = true
			outList = append(outList, v)
		}
	}
	return outList
}
/*
func main() {
	s1 := Set{elems: []int{1, 1, 2, 3, 4, 2, 1}}
	fmt.Println(s1.Elements())
}
*/
