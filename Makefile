ccwc:
	@echo "generating ccwc binary"
	@go build -o ./bin/ccwc ./cmd/ccwc/main.go
	@echo "ccwc binary is in ./bin"
