# Map or Array

```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/protovalidate
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkProtocGenValidate
BenchmarkProtocGenValidate-8            178371715                6.827 ns/op
BenchmarkProtocGenValidate-16           174528217                7.390 ns/op
BenchmarkProtocGenValidate-32           134606652                8.017 ns/op
BenchmarkProtoValidateGo
BenchmarkProtoValidateGo-8               2074315               590.8 ns/op
BenchmarkProtoValidateGo-16              1710081               649.2 ns/op
BenchmarkProtoValidateGo-32              1834952               596.5 ns/op
PASS
ok      performance-go-base/protovalidate       12.662s

```
