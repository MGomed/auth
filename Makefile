LOCAL_BIN:=$(CURDIR)/bin
API_PROTO:=user_api_v1
API:=user_api

BUILD_DIR:=build

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

.PHONY: test
test:
	go clean -testcache
	go test -coverprofile out/coverage/cover.out ./...
	go tool cover -html=cover.out -o out/coverage/coverage.html

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api
	make generate_mocks

generate_mocks:
	./generate.sh

generate-user-api:
	mkdir -p pkg/$(API)
	protoc --proto_path api/$(API_PROTO) \
	--go_out=pkg/$(API) --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/$(API) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/$(API_PROTO)/*

db-up:
	docker compose -f ${BUILD_DIR}/docker-compose.yml up --build -d pg migrator

db-down:
	docker compose -f ${BUILD_DIR}/docker-compose.yml down pg migrator
