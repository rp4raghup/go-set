package go-set

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
