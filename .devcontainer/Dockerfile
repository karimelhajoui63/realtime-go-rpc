FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

# Dev deps for proto files generation
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

# Dev deps for linter
RUN go install mvdan.cc/gofumpt@latest
