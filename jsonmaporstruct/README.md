# Proto Validate Benchmark

```shell
protoc --go_out=. --go_opt=paths=source_relative *.proto
```

```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/jsonmaporstruct
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkUnmarshalToObject
BenchmarkUnmarshalToObject-8              310347              3798 ns/op
BenchmarkUnmarshalToObject-16             307251              3858 ns/op
BenchmarkUnmarshalToObject-32             301804              3856 ns/op
BenchmarkUnmarshalToAny
BenchmarkUnmarshalToAny-8                 170671              7036 ns/op
BenchmarkUnmarshalToAny-16                170426              7296 ns/op
BenchmarkUnmarshalToAny-32                172196              7160 ns/op
BenchmarkUnmarshalWithJsoniter
BenchmarkUnmarshalWithJsoniter-8          707659              1685 ns/op
BenchmarkUnmarshalWithJsoniter-16         692643              1715 ns/op
BenchmarkUnmarshalWithJsoniter-32         667081              1734 ns/op
BenchmarkUnmarshalWithSonic
BenchmarkUnmarshalWithSonic-8            1441142               830.2 ns/op
BenchmarkUnmarshalWithSonic-16           1408960               858.5 ns/op
BenchmarkUnmarshalWithSonic-32           1328851               895.6 ns/op
BenchmarkUnmarshalWithEasyjson
BenchmarkUnmarshalWithEasyjson-8        738442906                1.622 ns/op
BenchmarkUnmarshalWithEasyjson-16       726652209                1.611 ns/op
BenchmarkUnmarshalWithEasyjson-32       741504214                1.620 ns/op
BenchmarkUnmarshalWithFastjson
BenchmarkUnmarshalWithFastjson-8         1716730               681.4 ns/op
BenchmarkUnmarshalWithFastjson-16        1756989               681.4 ns/op
BenchmarkUnmarshalWithFastjson-32        1729933               680.5 ns/op
BenchmarkUnmarshalWithProtojson
BenchmarkUnmarshalWithProtojson-8         134277              8677 ns/op
BenchmarkUnmarshalWithProtojson-16        138122              8661 ns/op
BenchmarkUnmarshalWithProtojson-32        136124              8820 ns/op
BenchmarkUnmarshalWithProto
BenchmarkUnmarshalWithProto-8            1000000              1080 ns/op
BenchmarkUnmarshalWithProto-16           1000000              1044 ns/op
BenchmarkUnmarshalWithProto-32           1000000              1044 ns/op
BenchmarkUnmarshalWithMsgpb
BenchmarkUnmarshalWithMsgpb-8            2063806               581.9 ns/op
BenchmarkUnmarshalWithMsgpb-16           2027316               588.0 ns/op
BenchmarkUnmarshalWithMsgpb-32           2003602               594.2 ns/op
BenchmarkUnmarshalWithGob
BenchmarkUnmarshalWithGob-8                57822             20915 ns/op
BenchmarkUnmarshalWithGob-16               57193             21168 ns/op
BenchmarkUnmarshalWithGob-32               56044             21793 ns/op
PASS
ok      performance-go-base/jsonmaporstruct     45.172s
```
