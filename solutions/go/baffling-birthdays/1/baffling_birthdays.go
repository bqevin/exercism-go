package bafflingbirthdays

import (
    "time"
    "math/rand/v2"
)

func SharedBirthday(dates []time.Time) bool {
	if len(dates) < 2 {
        return false;
    }
    for i, v := range dates {
        for j, k := range dates {
            if i == j { continue;} // skips current
            if v == k {
                return true;
            }
            if  k.Day() == v.Day() && k.Month() == v.Month() {
                return true;
            }
        }
    }

    return false;
}

func RandomBirthdates(size int) []time.Time {
    dates := make([]time.Time, size);
	for i := 0; i < size; i++ {
        min := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC).Unix();
        max := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC).Unix();
        // continue loop until valid year found (non-leap)
        for {
            randomSecs := rand.Int64N(max - min) + min;
            t := time.Unix(randomSecs, 0).UTC();
            if t.Year() % 4 != 0 && (t.Year() % 100 == 0 || t.Year() % 400 != 0){
                dates[i] = time.Unix(randomSecs, 0).UTC();
                break; // break loop
            }
        }
    }
    return dates;
}

func EstimatedProbability(size int) float64 {
    //P(shared) = 1 - P(not shared)
    if size >= 365 {
        return 1.0
    }
    noMatch := 1.0;
    for i := 0; i < size; i++ {
        // assuming not a leap year
        noMatch *= float64(365 - i) / 365.0
    }

	return (1 - noMatch) * 100;
}
