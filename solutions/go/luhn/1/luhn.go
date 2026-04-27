package luhn
import (
    "strings"
    "strconv"
)

func Valid(id string) bool {
    n := strings.TrimSpace(strings.ReplaceAll(id, " ", ""))
	l := len(n)
	sum := 0
	steps := 1

	if l < 2 {
		return false
	}

	for i := l - 1; i >= 0; i-- {
		curr, e := strconv.Atoi(string(n[i]))
        
        if e != nil {
            return false
        }
        
		if steps%2 == 0 {
			if curr*2 > 9 {
				sum += (curr * 2) - 9
			} else {
				sum += curr * 2
			}
		} else {
			sum += curr
		}
		steps++
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
