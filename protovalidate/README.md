# Proto Validate Benchmark

```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/protovalidate
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkProtocGenValidate
BenchmarkProtocGenValidate-8            166471131                6.897 ns/op
BenchmarkProtocGenValidate-16           172695788                7.071 ns/op
BenchmarkProtocGenValidate-32           171434694                7.012 ns/op
BenchmarkProtoValidateGo
BenchmarkProtoValidateGo-8               2251598               551.6 ns/op
BenchmarkProtoValidateGo-16              2259184               527.9 ns/op
BenchmarkProtoValidateGo-32              2237440               539.7 ns/op
BenchmarkProtoValidateAll
BenchmarkProtoValidateAll/ProtoValidateGo
BenchmarkProtoValidateAll/ProtoValidateGo-8              8275581               142.1 ns/op           176 B/op          6 allocs/op
BenchmarkProtoValidateAll/ProtoValidateGo-16            10811161               120.5 ns/op           176 B/op          6 allocs/op
BenchmarkProtoValidateAll/ProtoValidateGo-32            11063128               101.5 ns/op           176 B/op          6 allocs/op
BenchmarkProtoValidateAll/ProtocGenValidate
BenchmarkProtoValidateAll/ProtocGenValidate-8           860760510                1.363 ns/op           0 B/op          0 allocs/op
BenchmarkProtoValidateAll/ProtocGenValidate-16          1000000000               1.090 ns/op           0 B/op          0 allocs/op
BenchmarkProtoValidateAll/ProtocGenValidate-32          1000000000               0.9717 ns/op          0 B/op          0 allocs/op
PASS
ok      performance-go-base/protovalidate       20.378s
```
