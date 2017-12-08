## Performant implementation of Hamming distance

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/hamming)](https://goreportcard.com/report/jancajthaml-go/hamming)

### Usage ###

```
import "github.com/jancajthaml-go/hamming"

hamming.Distance("aba", "bba")
```

### Performance ###

```
BenchmarkHammingParallel-4    300000000           4.25 ns/op
BenchmarkHammingSerial-4      200000000           8.77 ns/op
```

test on your own by running `make benchmark`

## Resources

* [Wikipedia - Hamming distance](https://en.wikipedia.org/wiki/Hamming_distance)
