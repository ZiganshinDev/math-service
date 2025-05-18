workdir = $(pwd)

make up:
	cd ./deployments && docker compose up mathbot

make down:
	cd ./deployments && docker compose down

make lint:
	docker run --rm -v .:/app -w /app golangci/golangci-lint:v1.61.0-alpine golangci-lint run -c ./build/ci/.golangci.yml -v
	
make tidy: vendor
	go mod tidy

make vendor:
	go mod vendor
