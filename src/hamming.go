//
// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Hamming_distance
//
package main

// Distance returns hamming distance between two strings
func Distance(a, b string) int {
	return calculateDistance([]byte(b), []byte(a))
}

func calculateDistance(a, b []byte) int {
	if len(a) != len(b) {
		panic("input slices are of different length")
	}

	var (
		i     int
		l     int = len(a)
		count int
	)

loop:
	if i == l {
		return count
	}
	if a[i] != b[i] {
		count++
	}
	i++
	goto loop
}
