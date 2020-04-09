grpc:
	@protoc -I protoresources/ protoresources/*.proto --go_out=plugins=grpc:protoresources
# go build command
build:
	@echo " >> building binaries"
	@go build -v -o resource cmd/main.go

# go run command
run: build
	@./resource

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

