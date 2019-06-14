#### 1.article
https://doc.oschina.net/grpc?t=60133


#### 2.install
```bash
brew install protobuf

go get -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go

protoc --go_out=plugins=grpc:. *.proto

```
