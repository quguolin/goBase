

### v1.0.0
- 通过压测 没有 error full channel 日志
```
goos: darwin
goarch: amd64
pkg: goBase/log
BenchmarkFLog_Write-4   	 3000000	       522 ns/op
PASS
```
- bytes.NewBuffer 待优化