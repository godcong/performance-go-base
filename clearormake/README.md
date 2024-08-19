# Map or Array

```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/clearormake
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkClear
BenchmarkClear-8           12466            122873 ns/op
BenchmarkClear-16          12724            126425 ns/op
BenchmarkClear-32          12898            126445 ns/op
BenchmarkMake
BenchmarkMake-8            10000            243783 ns/op
BenchmarkMake-16           10000            237203 ns/op
BenchmarkMake-32           10000            236603 ns/op
PASS
ok      performance-go-base/clearormake 15.179s
```
