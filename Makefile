.PHONY: run build clean

BINARY = cmd/app/main.go

run:
	go run $(BINARY)
