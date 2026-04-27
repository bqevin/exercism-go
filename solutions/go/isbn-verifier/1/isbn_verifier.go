package isbnverifier

import (
    "strings"
    "strconv"
)
func IsValidISBN(isbn string) bool {
	s := strings.ReplaceAll(isbn, "-", "")
	l := len(s)
	step := 10
	sum := 0
	if l == 10 {
		for i := 0; i < l; i++ {
			curr, e := strconv.Atoi(string(s[i]))
			if e != nil {
				if s[i] == 'X' && i == 9 {
					curr = 10
				} else {
					return false
				}
			}
			sum += curr * step
			step -= 1
		}
		return sum%11 == 0
	}
	return false
}
