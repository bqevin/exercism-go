package differenceofsquares

func SquareOfSum(n int) int {
    count := 0
	for i := 1; i <= n; i++ {
         count += i
    }
    
    return count * count
}

func SumOfSquares(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
         count += i * i
    }
    
    return count
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
