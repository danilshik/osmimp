.PHONY: gen
gen:

	protoc --proto_path=./parser/pbf/internal/proto/ --go_out=./parser/pbf/internal/bin/ ./parser/pbf/internal/proto/*.proto


.PHONY: deps
deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
