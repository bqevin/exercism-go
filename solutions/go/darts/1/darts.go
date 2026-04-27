package darts

import "math"

func Score(x, y float64) int {
	r := math.Sqrt((x * x) + (y * y))

    if r > 10 {
        return 0
    }
    
    if r < 11 && r > 5 {
        return 1
    }
    
    if r < 6 && r > 1 {
        return 5
    }
    
    if r < 2 && r > -1 {
        return 10
    }

    return 0;
}
