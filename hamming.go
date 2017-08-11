package main

/**
 * Hamming distance algorithm
 *
 * @see https://en.wikipedia.org/wiki/Hamming_distance
 *
 * @author jan.cajthaml
 */
func HammingDistance(a, b []byte) int {
	if len(a) != len(b) {
		panic("input sets are of different length")
	}

	var (
		i     int = 0
		l     int = len(a)
		count int = 0
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
