LOCAL_BIN:=$(CURDIR)/bin
API_PROTO:=user_api_v1
API:=user_api

BUILD_DIR:=build

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

.PHONY: test
test:
	mkdir -p out/coverage
	go clean -testcache
	go test -coverprofile out/coverage/cover.out ./...
	go tool cover -html=out/coverage/cover.out -o out/coverage/coverage.html

install-deps:
	[ -f $(LOCAL_BIN)/golangci-lint ] || GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	[ -f $(LOCAL_BIN)/protoc-gen-go ] || GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	[ -f $(LOCAL_BIN)/protoc-gen-go-grpc ] || GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	[ -f $(LOCAL_BIN)/goose ] || GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	[ -f $(LOCAL_BIN)/mockgen ] || GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@v1.6.0
	[ -f $(LOCAL_BIN)/protoc-gen-validate ] || GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.0.4
	[ -f $(LOCAL_BIN)/protoc-gen-grpc-gateway ] || GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.20.0
	[ -f $(LOCAL_BIN)/protoc-gen-openapiv2 ] || GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.20.0
	[ -f $(LOCAL_BIN)/statik ] || GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7

get-deps:
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go get -u github.com/envoyproxy/protoc-gen-validate/validate
	GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway/v2/runtime

generate:
	make generate-user-api
	make generate_mocks
	make generate_statik

generate_mocks:
	./generate.sh

generate_statik:
	$(LOCAL_BIN)/pkg/statik -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'

generate-user-api:
	mkdir -p pkg/$(API)
	protoc --proto_path api/$(API_PROTO) --proto_path vendor.protogen \
	--go_out=pkg/$(API) --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/$(API) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/$(API) --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--grpc-gateway_out=pkg/$(API) --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/$(API_PROTO)/*

vendor-proto:
	@if [ ! -d vendor.protogen/validate ]; then \
		mkdir -p vendor.protogen/validate && \
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate && \
		mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate && \
		rm -rf vendor.protogen/protoc-gen-validate ; \
	fi
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis && \
		mkdir -p  vendor.protogen/google/ && \
		mv vendor.protogen/googleapis/google/api vendor.protogen/google && \
		rm -rf vendor.protogen/googleapis ; \
	fi
	@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
		mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
		git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
		mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
		rm -rf vendor.protogen/openapiv2 ;\
	fi

db-up:
	docker compose -f ${BUILD_DIR}/docker-compose.yml up --build -d

db-down:
	docker compose -f ${BUILD_DIR}/docker-compose.yml down
