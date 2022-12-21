GOPATH:=$(shell go env GOPATH)


.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: proto
proto:
	@protoc --proto_path=. --go-grpc_out=. --go_out=:. pb/*.proto

.PHONY: server
server:
	@go run server.go


.PHONY: client
client:
	@go run client.go

.PHONY: exit
exit:
	@go run consul-exit.go