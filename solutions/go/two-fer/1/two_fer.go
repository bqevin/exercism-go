package twofer

import "fmt"
// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
    n := name
	if n == "" {
		n = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", n)
}
