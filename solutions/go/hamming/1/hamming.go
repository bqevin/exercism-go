package hamming

import "errors"
func Distance(a, b string) (int, error) {
    lA := len(a)
	lB := len(b)
	e := errors.New("Invalid input")
	count := 0

	if lA != lB {
		return 0, e
	}

	for i := 0; i < lA; i++ {
		if a[i] != b[i] {
			count += 1
		}
	}

	return count, nil
}
