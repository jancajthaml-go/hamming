package hamming

import "testing"

func TestCalculations(t *testing.T) {
	expectDistance := func(a string, b string, c int) {
		d, err := Distance(a, b)
		if err != nil {
			t.Errorf("expected no error but got %+v", err)
			return
		}
		if d != c {
			t.Errorf("a: `" + a + "`, b: `" + b + "` expected " + string(rune(c+48)) + " , got " + string(rune(d+48)))
		}
	}

	// matches
	expectDistance("", "", 0)
	expectDistance("a", "a", 0)
	expectDistance("abc", "abc", 0)

	// swaps only
	expectDistance("a", "b", 1)
	expectDistance("ab", "ac", 1)
	expectDistance("ac", "bc", 1)
	expectDistance("abc", "axc", 1)
	expectDistance("xabxcdxxefxgx", "1ab2cd34ef5g6", 6)
}

func BenchmarkHammingSerial(b *testing.B) {
	left := "xxxx"
	right := "xxxy"

	for n := 0; n < b.N; n++ {
		Distance(left, right)
	}
}

func BenchmarkHammingParallel(b *testing.B) {
	left := "xxxx"
	right := "xxxy"

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Distance(left, right)
		}
	})
}
