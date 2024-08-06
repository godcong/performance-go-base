# Channel or Mutex

```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/channelormutex
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkNewChannel
BenchmarkNewChannel-8            2221848               550.1 ns/op
BenchmarkNewChannel-16           2414916               515.0 ns/op
BenchmarkNewChannel-32           2586192               479.1 ns/op
BenchmarkMutex
BenchmarkMutex-8                88250217                11.96 ns/op
BenchmarkMutex-16               92085976                12.15 ns/op
BenchmarkMutex-32               98747551                12.43 ns/op
PASS
ok      performance-go-base/channelormutex      11.202s
```