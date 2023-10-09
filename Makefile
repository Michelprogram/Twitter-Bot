clear:
	rm twitter-bot || true
	go clean -cache ./...
build:
	@go build -o twitter-bot -v ./...
run-preview:
	./twitter-bot
run:
	go run cmd/main.go
build-run: clear build run-preview
build-github:
	@go build -o twitter-bot -v ./...
	./twitter-bot 
