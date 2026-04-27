package phonenumber

import (
    "errors"
    "fmt"
    "strings"
    "strconv"
)

func Number(phoneNumber string) (string, error) {
	e := errors.New("invalid phone number")
	p := strings.NewReplacer("-", "", ".", "", " ", "", "(", "", ")", "").Replace(phoneNumber)
	p = strings.TrimSpace(p)
	l := len(p)

	if l >= 10 && l < 13 {
		if (l == 11 && p[0] != '1') ||
			l == 12 && p[0] != '+' {
			return p, e
		}
		// NXX NXX-XXXX
		// N is any digit from 2 through 9
		if p[l-7] < '2' || p[l-10] < '2' {
			return p, e
		}
		for _, v := range p[l-10:] {
			_, sE := strconv.Atoi(string(v))
			if sE != nil {
				return p, e
			}
		}

		return p[l-10:], nil
	}

	return p, e
}

func AreaCode(phoneNumber string) (string, error) {
	n, e := Number(phoneNumber)
	if e != nil {
		return n, e
	}

	return fmt.Sprintf("%s", n[:3]), nil
}

func Format(phoneNumber string) (string, error) {
	n, e := Number(phoneNumber)
	if e != nil {
		return n, e
	}

	return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:]), nil
}
