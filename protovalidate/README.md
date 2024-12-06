# Map or Array

```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/protovalidate
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkProtocGenValidate
BenchmarkProtocGenValidate-8            175539582                6.706 ns/op
BenchmarkProtocGenValidate-16           184073895                6.719 ns/op
BenchmarkProtocGenValidate-32           180479190                6.643 ns/op
BenchmarkProtoValidateGo
BenchmarkProtoValidateGo-8               2220597               536.1 ns/op
BenchmarkProtoValidateGo-16              2252672               536.8 ns/op
BenchmarkProtoValidateGo-32              2164675               549.1 ns/op
PASS
ok      performance-go-base/protovalidate       10.985s
```
