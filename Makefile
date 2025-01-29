APP_NAME=initiator

# Build the binary
build:
	@go build -o bin/$(APP_NAME) ./main.go

# Build the binary & run the tests & run the binary
start:
	@go build -o bin/$(APP_NAME) ./main.go
	@go test -v ./...
	@./bin/$(APP_NAME)

# Run the tests
test:
	@go test -v ./...