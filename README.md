## hamming distance

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/hamming)](https://goreportcard.com/report/jancajthaml-go/hamming)

### Usage ###

```
import "github.com/jancajthaml-go/hamming"

hamming.Distance("aba", "bba")
```

### Performance ###

```
BenchmarkHammingParallel-4  100000000  11.2 ns/op  0 B/op  0 allocs/op
BenchmarkHammingSerial-4    50000000   21.9 ns/op  0 B/op  0 allocs/op
```

verify your performance by running `make benchmark`

## Resources

* [Wikipedia - Hamming distance](https://en.wikipedia.org/wiki/Hamming_distance)
