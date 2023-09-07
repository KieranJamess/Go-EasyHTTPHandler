build:
	@go build -o ./bin/baseRoutingExample

run: build
	@./bin/baseRoutingExample

# test:
# 	@go test -v ./...