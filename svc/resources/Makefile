grpc:
	@protoc -I protoresources/ protoresources/*.proto --go_out=plugins=grpc:protoresources
# go build command
build-grpc:
	@echo " >> building binaries"
	@go build -v -o resource_grpc cmd/grpc/main.go
build-rest:
	@echo " >> building binaries"
	@go build -v -o resource_rest cmd/rest/main.go

# go run command
run-grpc: build-grpc
	@./resource_grpc
run-rest: build-rest
	@./resource_rest

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
migration-reset:
	@go run migrations/main.go reset

