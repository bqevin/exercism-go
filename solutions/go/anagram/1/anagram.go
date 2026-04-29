package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	var results []string
	s := []rune(strings.ToLower(subject))
	sL := len(s)
	for _, v := range candidates {
		c := []rune(strings.ToLower(v))
		sumS, sumC := 0, 0
		tS, tC := make(map[rune]int), make(map[rune]int)
		cL := len(c)
		// skip if subject is exactly similar to candidate
		if strings.ToLower(subject) == strings.ToLower(v) {
			continue
		}
		// skip when rune length mismatches
		if sL != cL {
			continue
		}
		// loop through subject & candidate runes
		for j := 0; j < sL; j++ {
			sumS += int(s[j])
			sumC += int(c[j])
			tS[s[j]]++
			tC[c[j]]++
			if j == sL-1 {
				if sumS == sumC && len(tS) == len(tC) {
					results = append(results, v)
				}
				clear(tC)
				clear(tS)
			}
		}
	}

	return results
}
