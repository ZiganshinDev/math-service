workdir = $(pwd)

make up: migrate-up
	cd ./deployments && docker compose up sindromitis

make down: migrate-down
	cd ./deployments && docker compose down

make migrate-up:
	cd ./deployments && docker compose up migrate-up

make migrate-down:
	cd ./deployments && docker compose up migrate-down

make lint:
	docker run --rm -v .:/app -w /app golangci/golangci-lint:v1.61.0-alpine golangci-lint run -c ./build/ci/.golangci.yml -v
	
make tidy: vendor
	go mod tidy

make vendor:
	go mod vendor

make protoc:
	protoc --proto_path=./proto --micro_out=./internal/proto --go_out=./internal/proto --go-grpc_out=./internal/proto author.proto
