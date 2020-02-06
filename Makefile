grpc:
	@protoc -I marketplaceproto/ marketplaceproto/*.proto --go_out=plugins=grpc:marketplaceproto

# go build command
build:
	@echo " >> building binaries"
	@go build -v -o marketplace main.go

# go run command
run: build
	@./marketplace

docker-run:
	@docker-compose up -d

docker-stop:
	@docker-compose down