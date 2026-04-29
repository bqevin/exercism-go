package armstrongnumbers

import (
    "math"
    "strconv"
    )

func IsNumber(n int) bool {
	s := strconv.Itoa(n)
	sum := 0
	for _, v := range s {
		i, _ := strconv.Atoi(string(v))
		sum += int(math.Pow(float64(i), float64(len(s))))
	}

	return sum == n
}
