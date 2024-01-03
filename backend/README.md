[Official documentation](https://connectrpc.com/docs/go/getting-started) of Connect


### Commands used in bulk:

```sh
cd backend

go mod init rpc-server

go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

go install google.golang.org/protobuf/cmd/protoc-gen-es@latest

buf lint ../proto
buf generate --template ../buf.gen.go.yaml ../proto

go run cmd/server/main.go

# Linter
go install mvdan.cc/gofumpt@latest
gofumpt -l -w .
# or for VSCode:
# "go.useLanguageServer": true,
# "gopls": {
# 	"formatting.gofumpt": true,
# },
```

### Bref explanation

A color is stored locally. It can be changed through the RPC API. 
From this API, you can also subscribe to the color's changes and receive a message from the server everytimes that the color change.
Technically, there is a RabbitMQ exchange that broadcast the color changes to every subscriber.

### Core architecture

I used the hexagonal architecture in order to be able to easly swap the handler and try the difference between RPC, SSE and WS.
See this repo for root directories: https://github.com/golang-standards/project-layout
In the `internal` directory, I have:
 - `core`: where the "buisness logic" lives (nothing fancy in this case though... there is just the logic to store a color and change it or subscribe to its changes)
 - `port`: where the interfaces are defined
    - `inbound`: implemented in `core/usecase/*.go` and used by the inbound adapter(s) in `adapter/inbound/*/*.go`
    - `outbound`: implemented in the outbound adapter(s) `adapter/outbound/*/*.go` and used by the usecases `core/usecase/*.go`
 - `adapter`: where the concret implementation of the interfaces (=port) are defined
