package clock

import "fmt"

type Clock struct {
    hrs, mins int
}

func New(h, m int) Clock {
    hrs := 0
	mins := 0

	hrs += h + m/60
	mins = m % 60

    for mins < 0 {
		mins += 60
		hrs -= 1
	}

	for hrs < 0 {
		hrs += 24
	}
    
	if hrs >= 24 {
		hrs %= 24
	}

	return Clock{
		hrs:  hrs,
		mins: mins,
	}
}

func (c Clock) Add(m int) Clock {
    c.mins += m
    c.hrs += c.mins / 60
	c.mins %= 60

	if c.hrs >= 24 {
		c.hrs %= 24
	}

	return c
}

func (c Clock) Subtract(m int) Clock {  
    c.mins -= m
	for c.mins < 0 {
		c.mins += 60
		c.hrs -= 1
	}
	
	for c.hrs < 0 {
		c.hrs += 24
	}

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hrs, c.mins)
}
