
## Run compiling command for PROTO Gcrp
```go
protoc --go_out=api-gateway/ --go_opt=paths=source_relative --go-grpc_out=api-gateway/ --go-grpc_opt=paths=source_relative protos/book.proto
```