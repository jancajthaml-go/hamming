//
// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Hamming_distance
//
package hamming

import "fmt"

// Distance returns hamming distance between two strings
func Distance(a string, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("input slices are of different length")
	}

	aa := []byte(a)
	ab := []byte(b)

	var (
		i     int
		l     int = len(aa)
		count int
	)

loop:
	if i == l {
		return count, nil
	}
	if aa[i] != ab[i] {
		count++
	}
	i++
	goto loop

}
