//
// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Hamming_distance
//
package main

func Distance(a, b string) int {
	return calculate_distance([]byte(b), []byte(a))
}

func calculate_distance(a, b []byte) int {
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
