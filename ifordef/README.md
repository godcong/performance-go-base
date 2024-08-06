# If or Default


## with pointer `(c *ConditionXXX) Call() bool`
```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/ifordef
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkDef
BenchmarkDef-8          892819350                1.328 ns/op
BenchmarkDef-16         860450674                1.281 ns/op
BenchmarkDef-32         948777104                1.288 ns/op
BenchmarkIf
BenchmarkIf-8           883647224                1.413 ns/op
BenchmarkIf-16          856830724                1.446 ns/op
BenchmarkIf-32          852986982                1.376 ns/op
PASS
ok      performance-go-base/ifordef     9.446s
```

# without pointer `(c ConditionXXX) Call() bool`
```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/ifordef
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkDef
BenchmarkDef-8          820832734                1.328 ns/op
BenchmarkDef-16         892290829                1.295 ns/op
BenchmarkDef-32         874222669                1.268 ns/op
BenchmarkIf
BenchmarkIf-8           847385814                1.470 ns/op
BenchmarkIf-16          857430708                1.369 ns/op
BenchmarkIf-32          869671089                1.394 ns/op
PASS
ok      performance-go-base/ifordef     7.896s
```