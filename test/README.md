## 覆盖率分析
- 覆盖率
```
go test -cover
```

- html展示覆盖率
```
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html 
```