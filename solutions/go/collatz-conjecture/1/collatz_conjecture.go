package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    steps := 0
    e := errors.New("Not valid number")
	if n < 1 {
        return 0, e;
    }
    if n == 1 {
        return 0, nil;
    }

    for n > 1 {
       if n % 2 == 0 {
           n = n / 2
       } else {
           n = (n * 3) + 1
       }
        steps++ 
    }

    return steps, nil
}
