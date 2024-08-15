# Map or Array

```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/reflect
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkCallDeep
BenchmarkCallDeep-8             26035112                49.88 ns/op
BenchmarkCallDeep-16            25533598                45.76 ns/op
BenchmarkCallDeep-32            25494324                43.55 ns/op
BenchmarkCallDeepArr
BenchmarkCallDeepArr-8          46053766                27.62 ns/op
BenchmarkCallDeepArr-16         46780342                29.10 ns/op
BenchmarkCallDeepArr-32         48930860                27.95 ns/op
BenchmarkCallDeepBig
BenchmarkCallDeepBig-8          44331808                24.03 ns/op
BenchmarkCallDeepBig-16         53265153                24.76 ns/op
BenchmarkCallDeepBig-32         43767666                24.36 ns/op
BenchmarkCallZero
BenchmarkCallZero-8             910194174                1.374 ns/op
BenchmarkCallZero-16            898171994                1.314 ns/op
BenchmarkCallZero-32            913341405                1.361 ns/op
BenchmarkCallZeroArr
BenchmarkCallZeroArr-8          194352410                6.525 ns/op
BenchmarkCallZeroArr-16         181980640                6.789 ns/op
BenchmarkCallZeroArr-32         181498304                6.635 ns/op
BenchmarkCallZeroBig
BenchmarkCallZeroBig-8          166375670                7.447 ns/op
BenchmarkCallZeroBig-16         168787884                6.959 ns/op
BenchmarkCallZeroBig-32         167411648                7.016 ns/op
BenchmarkCallPreZero
BenchmarkCallPreZero-8          722398507                1.609 ns/op
BenchmarkCallPreZero-16         764583474                1.554 ns/op
BenchmarkCallPreZero-32         794697776                1.609 ns/op
BenchmarkCallPreZeroArr
BenchmarkCallPreZeroArr-8       179006317                6.648 ns/op
BenchmarkCallPreZeroArr-16      184260526                7.049 ns/op
BenchmarkCallPreZeroArr-32      179789097                6.450 ns/op
BenchmarkCallPreZeroBig
BenchmarkCallPreZeroBig-8       176909870                6.557 ns/op
BenchmarkCallPreZeroBig-16      168453915                6.657 ns/op
BenchmarkCallPreZeroBig-32      179813964                6.433 ns/op
PASS
ok      performance-go-base/reflect     42.674s
```
