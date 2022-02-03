## Generate `PROTOBUF`

```bash
export PATH="$PATH:$(go env GOPATH)/bin"

protoc --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    proto/app.proto
```

## Generat **SWAGGER**

```bash
protoc --proto_path=. \
    --swagger_out=logtostderr=true:. \
    app.proto
```