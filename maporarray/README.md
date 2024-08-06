# Map or Array

```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/maporarray
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkMap
BenchmarkMap-8          34078516                44.75 ns/op
BenchmarkMap-16         42020042                44.69 ns/op
BenchmarkMap-32         45079564                45.14 ns/op
BenchmarkArrKey
BenchmarkArrKey-8       1000000000               0.2430 ns/op
BenchmarkArrKey-16      1000000000               0.2379 ns/op
BenchmarkArrKey-32      1000000000               0.2347 ns/op
BenchmarkArrVal
BenchmarkArrVal-8         991227            130794 ns/op
BenchmarkArrVal-16       1000000            132119 ns/op
BenchmarkArrVal-32        990074            135975 ns/op
PASS
ok      performance-go-base/maporarray  411.684s
```
