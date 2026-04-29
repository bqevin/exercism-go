package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	results := make(map[string]int)
    for i, v := range in {
        for _, k := range v {
            results[strings.ToLower(k)] = i
        }
    }
    
    return results;
}
