GOPATH:=$(shell go env GOPATH)


.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go mod tidy
.PHONY: update
update:
	@go get -u


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