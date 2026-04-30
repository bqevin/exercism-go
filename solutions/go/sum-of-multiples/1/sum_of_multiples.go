package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	res := map[int]int{}
	sum := 0
	for i := 1; i < limit; i++ {
		temp := 0
		for _, d := range divisors {
            if d == 0 {
				continue
			}
			if i%d == 0 {
				res[i]++
				if res[i] <= 1 {
					temp = i
				}
			}
		}
		if temp > 0 {
			sum += temp
		}

		if i == limit {
			clear(res)
		}
	}
	return sum
}
