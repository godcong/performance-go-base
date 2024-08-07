# Map or Array

```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/maporarray
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkMap
BenchmarkMap-8           8531034               168.7 ns/op
BenchmarkMap-16          9560685               169.2 ns/op
BenchmarkMap-32          9876852               168.7 ns/op
BenchmarkArrKey
BenchmarkArrKey-8       1000000000               0.2653 ns/op
BenchmarkArrKey-16      1000000000               0.2483 ns/op
BenchmarkArrKey-32      1000000000               0.2598 ns/op
BenchmarkArrVal
BenchmarkArrVal-8       64203739                68.00 ns/op
BenchmarkArrVal-16      53901091                28.10 ns/op
BenchmarkArrVal-32      58873166                21.94 ns/op
PASS
ok      performance-go-base/maporarray  24.595s
```
