clear:
	rm twitter-bot
	go clean -cache ./...
build:
	@go build -o twitter-bot -v ./...
run-preview:
	./twitter-bot
run:
	go run cmd/main.go
build-run: clear build run-preview
