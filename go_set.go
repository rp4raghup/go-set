package go_set

type Set struct {
	elems    []int
	elemsMap map[int]bool
}

func NewSet(elems []int) *Set {
	s := Set{}
	s.elemsMap = make(map[int]bool)
	for _, e := range elems {
		if _, ok := s.elemsMap[e]; !ok {
			s.elemsMap[e] = true
			s.elems = append(s.elems, e)
		}
	}
	return &s
}

func (s Set) Val() []int {
	return s.elems
}

func (s *Set) Add(e int) {
	if _, ok := s.elemsMap[e]; !ok {
		s.elemsMap[e] = true
		s.elems = append(s.elems, e)
	}
}

func (s1 *Set) Update(s2 Set) {
	for _, e := range s2.Val() {
		if _, ok := s1.elemsMap[e]; !ok {
			s1.elemsMap[e] = true
			s1.elems = append(s1.elems, e)
		}
	}
}

func (s1 Set) Equals(s2 Set) bool {
	if len(s1.Val()) != len(s2.Val()) {
		return false
	}

	for _, e := range s2.Val() {
		if _, ok := s1.elemsMap[e]; !ok {
			return false
		}
	}
	return true
}
