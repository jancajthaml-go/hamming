//
// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Hamming_distance
//
package hamming

import "fmt"

// Distance returns hamming distance between two strings
func Distance(left string, right string) (int, error) {
	if len(left) != len(right) {
		return 0, fmt.Errorf("inputs are of different length")
	}
	var (
		i     int
		l     int = len(left)
		count int
	)
loop:
	if i == l {
		return count, nil
	}
	if left[i] != right[i] {
		count++
	}
	i++
	goto loop
}
