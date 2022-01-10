//go:build tools

package tools

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-lint"
	_ "github.com/go-bridget/twirp-swagger-gen/cmd/protoc-gen-twirp-swagger"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/twitchtv/twirp/protoc-gen-twirp"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
