# Libp2p or TCP

```
>go test -v -bench='$' -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/libp2portcp
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkLibp2pTransferOneStream
BenchmarkLibp2pTransferOneStream-8                         19347             59017 ns/op
BenchmarkLibp2pTransferOneStream-16                        17774             65926 ns/op
BenchmarkLibp2pTransferOneStream-32                        17054             80660 ns/op
BenchmarkLibp2pTransferMultiConnectStream
BenchmarkLibp2pTransferMultiConnectStream-8                 4707            249779 ns/op
BenchmarkLibp2pTransferMultiConnectStream-16                4444            260863 ns/op
BenchmarkLibp2pTransferMultiConnectStream-32                5396            203863 ns/op
BenchmarkTCPTransferOneConnect
BenchmarkTCPTransferOneConnect-8                           73524             14596 ns/op
BenchmarkTCPTransferOneConnect-16                          78642             15330 ns/op
BenchmarkTCPTransferOneConnect-32                          77779             15566 ns/op
BenchmarkTCPTransferMultiConnect
BenchmarkTCPTransferMultiConnect-8                          6344            291619 ns/op
BenchmarkTCPTransferMultiConnect-16                         4200            269270 ns/op
BenchmarkTCPTransferMultiConnect-32                         4258            340425 ns/op
PASS
ok      performance-go-base/libp2portcp 27.263s
```
