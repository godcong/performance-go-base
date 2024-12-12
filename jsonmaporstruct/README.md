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
BenchmarkUnmarshalToObject-8                              168520              6761 ns/op
BenchmarkUnmarshalToObject-16                             176020              6653 ns/op
BenchmarkUnmarshalToObject-32                             182457              6690 ns/op
BenchmarkUnmarshalToObjectMarshal
BenchmarkUnmarshalToObjectMarshal-8                       151572              7948 ns/op
BenchmarkUnmarshalToObjectMarshal-16                      144523              7971 ns/op
BenchmarkUnmarshalToObjectMarshal-32                      149988              7963 ns/op
BenchmarkUnmarshalToAny
BenchmarkUnmarshalToAny-8                                 174324              6871 ns/op
BenchmarkUnmarshalToAny-16                                171928              6992 ns/op
BenchmarkUnmarshalToAny-32                                175054              6980 ns/op
BenchmarkUnmarshalToAnyMarshal
BenchmarkUnmarshalToAnyMarshal-8                           91300             13362 ns/op
BenchmarkUnmarshalToAnyMarshal-16                          91141             13047 ns/op
BenchmarkUnmarshalToAnyMarshal-32                          91306             13792 ns/op
BenchmarkUnmarshalWithJsoniter
BenchmarkUnmarshalWithJsoniter-8                          593527              1976 ns/op
BenchmarkUnmarshalWithJsoniter-16                         628212              2007 ns/op
BenchmarkUnmarshalWithJsoniter-32                         605928              2019 ns/op
BenchmarkUnmarshalWithJsoniterMarshal
BenchmarkUnmarshalWithJsoniterMarshal-8                   395848              3053 ns/op
BenchmarkUnmarshalWithJsoniterMarshal-16                  383648              3084 ns/op
BenchmarkUnmarshalWithJsoniterMarshal-32                  392646              3124 ns/op
BenchmarkUnmarshalWithSonic
BenchmarkUnmarshalWithSonic-8                            1000000              1167 ns/op
BenchmarkUnmarshalWithSonic-16                            980367              1199 ns/op
BenchmarkUnmarshalWithSonic-32                            955124              1240 ns/op
BenchmarkUnmarshalWithSonicMarshal
BenchmarkUnmarshalWithSonicMarshal-8                      605100              1947 ns/op
BenchmarkUnmarshalWithSonicMarshal-16                     566449              1958 ns/op
BenchmarkUnmarshalWithSonicMarshal-32                     587670              1994 ns/op
BenchmarkUnmarshalWithFastjson
BenchmarkUnmarshalWithFastjson-8                          379670              3273 ns/op
BenchmarkUnmarshalWithFastjson-16                         373647              3217 ns/op
BenchmarkUnmarshalWithFastjson-32                         336097              3558 ns/op
BenchmarkUnmarshalWithFastjsonMarshal
BenchmarkUnmarshalWithFastjsonMarshal-8                   297974              4092 ns/op
BenchmarkUnmarshalWithFastjsonMarshal-16                  299167              4274 ns/op
BenchmarkUnmarshalWithFastjsonMarshal-32                  258829              4499 ns/op
BenchmarkUnmarshalWithProtojson
BenchmarkUnmarshalWithProtojson-8                         129660              8861 ns/op
BenchmarkUnmarshalWithProtojson-16                        130806              8844 ns/op
BenchmarkUnmarshalWithProtojson-32                        135177              8763 ns/op
BenchmarkUnmarshalWithProtojsonMarshal
BenchmarkUnmarshalWithProtojsonMarshal-8                   82230             14611 ns/op
BenchmarkUnmarshalWithProtojsonMarshal-16                  78631             15017 ns/op
BenchmarkUnmarshalWithProtojsonMarshal-32                  64330             15689 ns/op
BenchmarkUnmarshalWithEasyjson
BenchmarkUnmarshalWithEasyjson-8                          481224              2605 ns/op
BenchmarkUnmarshalWithEasyjson-16                         457014              2632 ns/op
BenchmarkUnmarshalWithEasyjson-32                         458424              2753 ns/op
BenchmarkUnmarshalWithEasyjsonMarshal
BenchmarkUnmarshalWithEasyjsonMarshal-8                   374362              3107 ns/op
BenchmarkUnmarshalWithEasyjsonMarshal-16                  392278              2984 ns/op
BenchmarkUnmarshalWithEasyjsonMarshal-32                  365586              3058 ns/op
BenchmarkUnmarshalWithEasyjsonSTD
BenchmarkUnmarshalWithEasyjsonSTD-8                       174400              6519 ns/op
BenchmarkUnmarshalWithEasyjsonSTD-16                      190734              6536 ns/op
BenchmarkUnmarshalWithEasyjsonSTD-32                      180848              6535 ns/op
BenchmarkUnmarshalWithEasyjsonSTDMarshal
BenchmarkUnmarshalWithEasyjsonSTDMarshal-8                112483             10048 ns/op
BenchmarkUnmarshalWithEasyjsonSTDMarshal-16               118179             10414 ns/op
BenchmarkUnmarshalWithEasyjsonSTDMarshal-32               119727             10166 ns/op
BenchmarkUnmarshalWithEasyjsonOneInstance
BenchmarkUnmarshalWithEasyjsonOneInstance-8             87226044                13.58 ns/op
BenchmarkUnmarshalWithEasyjsonOneInstance-16            91510843                13.85 ns/op
BenchmarkUnmarshalWithEasyjsonOneInstance-32            99721610                13.21 ns/op
BenchmarkUnmarshalWithEasyjsonOneInstanceMarshal
BenchmarkUnmarshalWithEasyjsonOneInstanceMarshal-8       1748872               674.2 ns/op
BenchmarkUnmarshalWithEasyjsonOneInstanceMarshal-16      1698405               732.0 ns/op
BenchmarkUnmarshalWithEasyjsonOneInstanceMarshal-32      1569370               803.9 ns/op
BenchmarkUnmarshalWithProto
BenchmarkUnmarshalWithProto-8                            1000000              1062 ns/op
BenchmarkUnmarshalWithProto-16                           1000000              1033 ns/op
BenchmarkUnmarshalWithProto-32                           1000000              1190 ns/op
BenchmarkUnmarshalWithProtoMarshal
BenchmarkUnmarshalWithProtoMarshal-8                      713758              1689 ns/op
BenchmarkUnmarshalWithProtoMarshal-16                     689064              1680 ns/op
BenchmarkUnmarshalWithProtoMarshal-32                     706422              1826 ns/op
BenchmarkUnmarshalWithMsgpb
BenchmarkUnmarshalWithMsgpb-8                            1000000              1113 ns/op
BenchmarkUnmarshalWithMsgpb-16                           1000000              1098 ns/op
BenchmarkUnmarshalWithMsgpb-32                           1000000              1147 ns/op
BenchmarkUnmarshalWithMsgpbMarshal
BenchmarkUnmarshalWithMsgpbMarshal-8                      532483              2265 ns/op
BenchmarkUnmarshalWithMsgpbMarshal-16                     545409              2257 ns/op
BenchmarkUnmarshalWithMsgpbMarshal-32                     525423              2353 ns/op
BenchmarkUnmarshalWithGob
BenchmarkUnmarshalWithGob-8                                51574             22607 ns/op
BenchmarkUnmarshalWithGob-16                               51985             22815 ns/op
BenchmarkUnmarshalWithGob-32                               52954             23353 ns/op
BenchmarkUnmarshalWithGobMarshal
BenchmarkUnmarshalWithGobMarshal-8                         37002             33061 ns/op
BenchmarkUnmarshalWithGobMarshal-16                        35649             31578 ns/op
BenchmarkUnmarshalWithGobMarshal-32                        36490             32144 ns/op
PASS
ok      performance-go-base/jsonmaporstruct     97.042s
```
