# Proto Validate Benchmark

```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/protovalidate
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkProtocGenValidate
BenchmarkProtocGenValidate-8            157742784                7.578 ns/op
BenchmarkProtocGenValidate-16           157807713                7.546 ns/op
BenchmarkProtocGenValidate-32           158929808                7.581 ns/op
BenchmarkProtoValidateGo
BenchmarkProtoValidateGo-8               2275038               527.7 ns/op
BenchmarkProtoValidateGo-16              2231749               546.2 ns/op
BenchmarkProtoValidateGo-32              1960032               539.3 ns/op
BenchmarkProtoValidateAll
BenchmarkProtoValidateAll/ProtoValidateGo
BenchmarkProtoValidateAll/ProtoValidateGo-8              7650406               145.1 ns/op           176 B/op          6 allocs/op
BenchmarkProtoValidateAll/ProtoValidateGo-16            10536250               122.9 ns/op           176 B/op          6 allocs/op
BenchmarkProtoValidateAll/ProtoValidateGo-32            11404188               102.9 ns/op           176 B/op          6 allocs/op
BenchmarkProtoValidateAll/ProtocGenValidate
BenchmarkProtoValidateAll/ProtocGenValidate-8           819668211                1.491 ns/op           0 B/op          0 allocs/op
BenchmarkProtoValidateAll/ProtocGenValidate-16          1000000000               1.140 ns/op           0 B/op          0 allocs/op
BenchmarkProtoValidateAll/ProtocGenValidate-32          1000000000               0.9852 ns/op          0 B/op          0 allocs/op
PASS
ok      performance-go-base/protovalidate       20.291s
```
