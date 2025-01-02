# Orm Benchmark

```
>go test -v -bench=^$ -cpu="8,16,32" .
goos: windows
goarch: amd64
pkg: performance-go-base/orm
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkORMCreate
BenchmarkORMCreate/GORM-Create
BenchmarkORMCreate/GORM-Create-8                   33870             32111 ns/op
BenchmarkORMCreate/GORM-Create-16                  38674             33893 ns/op
BenchmarkORMCreate/GORM-Create-32                  37792             31385 ns/op
BenchmarkORMCreate/XORM-Create
BenchmarkORMCreate/XORM-Create-8                  139711              8796 ns/op
BenchmarkORMCreate/XORM-Create-16                 135826              8904 ns/op
BenchmarkORMCreate/XORM-Create-32                 133748              9420 ns/op
BenchmarkORMCreate/Ent-Create
BenchmarkORMCreate/Ent-Create-8                   122596             10310 ns/op
BenchmarkORMCreate/Ent-Create-16                  120536             10238 ns/op
BenchmarkORMCreate/Ent-Create-32                   99466             11797 ns/op
BenchmarkORMQuery
    orm_test.go:33: Start generating test data...
    orm_test.go:116: The test data is generated
BenchmarkORMQuery/GORM-SimpleQuery
BenchmarkORMQuery/GORM-SimpleQuery-8               70246             14573 ns/op
BenchmarkORMQuery/GORM-SimpleQuery-16              86148             15355 ns/op
BenchmarkORMQuery/GORM-SimpleQuery-32              78542             16012 ns/op
BenchmarkORMQuery/XORM-SimpleQuery
BenchmarkORMQuery/XORM-SimpleQuery-8               90463             12772 ns/op
BenchmarkORMQuery/XORM-SimpleQuery-16             102720             11883 ns/op
BenchmarkORMQuery/XORM-SimpleQuery-32             101209             12542 ns/op
BenchmarkORMQuery/Ent-SimpleQuery
BenchmarkORMQuery/Ent-SimpleQuery-8                81667             14749 ns/op
BenchmarkORMQuery/Ent-SimpleQuery-16               81393             14660 ns/op
BenchmarkORMQuery/Ent-SimpleQuery-32               80894             14957 ns/op
BenchmarkORMDelete
    orm_test.go:33: Start generating test data...
    orm_test.go:116: The test data is generated
BenchmarkORMDelete/GORM-SimpleDelete
BenchmarkORMDelete/GORM-SimpleDelete-8             68787             16590 ns/op
BenchmarkORMDelete/GORM-SimpleDelete-16            69205             16666 ns/op
BenchmarkORMDelete/GORM-SimpleDelete-32            69115             17054 ns/op
BenchmarkORMDelete/XORM-SimpleDelete
BenchmarkORMDelete/XORM-SimpleDelete-8            153218              7617 ns/op
BenchmarkORMDelete/XORM-SimpleDelete-16           147094              8032 ns/op
BenchmarkORMDelete/XORM-SimpleDelete-32           136602              8327 ns/op
BenchmarkORMDelete/Ent-SimpleDelete
BenchmarkORMDelete/Ent-SimpleDelete-8             245234              4865 ns/op
BenchmarkORMDelete/Ent-SimpleDelete-16            244687              4859 ns/op
BenchmarkORMDelete/Ent-SimpleDelete-32            249050              5178 ns/op
PASS
ok      performance-go-base/orm 40.546s
```
