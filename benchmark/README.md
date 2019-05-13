##### command
> 执行所有函数 ``go test -bench=".*"``

> 执行指定函数 ``go test -bench="BenchmarkFib10"``

> 生成内存和cpu分析文件 ``go test -bench="BenchmarkFib10" -cpuprofile cpu.out -memprofile mem.out``

> 分析文件 ``go tool pprof cpu.out``