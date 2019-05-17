### cpu debug info
    go test -cpuprofile cpu.prof
    go tool pprof cpu.prof
    - top
    - web
    
### mem debug info
    go test -memprofile mem.out
    go tool pprof mem.prof
    - top
    - web    

### all info
     go test -bench=. -memprofile=mem.out  -cpuprofile=cpu.out    
     
### Result info
```
goos: darwin
   goarch: amd64
   pkg: goBase/log
   BenchmarkFLog_Write-4   	 3000000	       502 ns/op
   PASS
   ok  	goBase/log	2.159s
```     
- **测试时间默认是1秒，也就是1秒的时间，调用三百万次，每次调用花费502纳秒**
- -8 运行时对应的GOMAXPROCS
- 3000000 表示运行for循环的次数
- 502 ns/op表示每次需要话费117纳秒

### Other
- go test -bench=. -benchtime=3s
