package set

// A set DS can easily be created using a map DS
// https://flaviocopes.com/golang-data-structure-set/
func main() {
	s1 := []int{1, 1, 1, 2, 3, 4, 5, 6}
	s2 := []int{1, 2, 3, 4, 6}

	s := map[int]bool{5: true, 2: true}
	_, ok := s[6] // check for existence
	s[8] = true   // add element
	delete(s, 2)  // remove element

	// UNION OPERATION
	s_union := map[int]bool{}
	for k, _ := range s1 {
		s_union[k] = true
	}
	for k, _ := range s2 {
		s_union[k] = true
	}

	// INTERSECTION
	s_intersection := map[int]bool{}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1 // better to iterate over a shorter set
	}
	for k, _ := range s1 {
		if s2[k] {
			s_intersection[k] = true
		}
	}
}
