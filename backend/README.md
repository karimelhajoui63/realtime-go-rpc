[Official documentation](https://connectrpc.com/docs/go/getting-started) of Connect


### Commands used:

go mod init rpc-server

go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

buf lint proto
buf generate

go run cmd/server/main.go
go run cmd/client/main.go --color=green