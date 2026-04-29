package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

// Foldl - Fold from Left to Right. Arguments are in form B(init), A(value)
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

// Foldr - Fold from Right to Left. Arguments are in form A(value), B(init)
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	res := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}
    
	return res
}

func (s IntList) Length() int {
	count := 0
	for range s {
		count += 1
	}
	return count
}

func (s IntList) Map(fn func(int) int) IntList {
	res := make(IntList, 0)
	for _, v := range s {
		res = append(res, fn(v))
	}
    
	return res
}

func (s IntList) Reverse() IntList {
	l := s.Length()
	for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	lA, lB := s.Length(), lst.Length()
	final := make(IntList, lA+lB)
	for i, v := range s {
		final[i] = v
	}
	for i, v := range lst {
		final[i+lA] = v
	}

	return final
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
		s = s.Append(l)
	}
	return s
}
