# Proto Validate Benchmark

```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/jsonmaporstruct
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkUnmarshalToObject
BenchmarkUnmarshalToObject-8              327976              3649 ns/op
BenchmarkUnmarshalToObject-16             332256              3730 ns/op
BenchmarkUnmarshalToObject-32             319563              3745 ns/op
BenchmarkUnmarshalToAny
BenchmarkUnmarshalToAny-8                 166041              7242 ns/op
BenchmarkUnmarshalToAny-16                169965              7127 ns/op
BenchmarkUnmarshalToAny-32                175149              7000 ns/op
BenchmarkUnmarshalWithJsoniter
BenchmarkUnmarshalWithJsoniter-8          823298              1502 ns/op
BenchmarkUnmarshalWithJsoniter-16         759200              1562 ns/op
BenchmarkUnmarshalWithJsoniter-32         758758              1563 ns/op
BenchmarkUnmarshalWithSonic
BenchmarkUnmarshalWithSonic-8            1689787               716.2 ns/op
BenchmarkUnmarshalWithSonic-16           1675312               713.0 ns/op
BenchmarkUnmarshalWithSonic-32           1624316               738.1 ns/op
BenchmarkUnmarshalWithEasyjson
BenchmarkUnmarshalWithEasyjson-8        734221128                1.655 ns/op
BenchmarkUnmarshalWithEasyjson-16       724256745                1.618 ns/op
BenchmarkUnmarshalWithEasyjson-32       755287484                1.668 ns/op
BenchmarkUnmarshalWithFastjson
BenchmarkUnmarshalWithFastjson-8         1788475               660.7 ns/op
BenchmarkUnmarshalWithFastjson-16        1810641               652.0 ns/op
BenchmarkUnmarshalWithFastjson-32        1826226               649.9 ns/op
BenchmarkUnmarshalWithProtojson
BenchmarkUnmarshalWithProtojson-8         134068              8958 ns/op
BenchmarkUnmarshalWithProtojson-16        138591              8603 ns/op
BenchmarkUnmarshalWithProtojson-32        135886              8847 ns/op
PASS
ok      performance-go-base/jsonmaporstruct     32.242s
```
