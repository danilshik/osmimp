
#GO_BIN=$(shell go env GOPATH)/bin
#BUF_EXE=$(GO_BIN)\buf$(shell go env GOEXE)


.PHONY: run
run:
	go run cmd/main.go

.PHONY: run with import parameters
run_with_import_parameters:
	go run cmd/main.go import

.PHONY: gen
gen:
	protoc --proto_path=./parser/pbf/internal/proto/ \
           --go_out=./parser/pbf/internal/model/ \
           --go_opt=paths=source_relative \
           --go_opt=M. \
           ./parser/pbf/internal/proto/*.proto



.PHONY: deps
deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
