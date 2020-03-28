grpc:
	@protoc -I catalogproto/ catalogproto/*.proto --go_out=plugins=grpc:catalogproto

# go build command
build:
	@echo " >> building binaries"
	@go build -v -o marketplace cmd/main.go

# go run command
run: build
	@./marketplace

docker-run:
	@docker-compose up -d

docker-stop:
	@docker-compose down

dep:
	@dep ensure -v
	@dep ensure -v --vendor-only

init:
	@dep init -v