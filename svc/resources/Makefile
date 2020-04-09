grpc:
	@protoc -I protoresources/ protoresources/*.proto --go_out=plugins=grpc:protoresources
# go build command
build-grpc:
	@echo " >> building binaries"
	@go build -v -o resource_grpc cmd/grpc/main.go

# go run command
run-grpc: build-grpc
	@./resource_grpc

docker-run:
	@docker-compose up -d

docker-stop:
	@docker-compose down

dep:
	@dep ensure -v
	@dep ensure -v --vendor-only

init:
	@dep init -v

migration-init:
	@go run migrations/main.go init

migration-up:
	@go run migrations/main.go up

migration-down:
	@go run migrations/main.go down

