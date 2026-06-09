watch/build:
	@go build -o ./tmp/server ./cmd/server

watch:
	@go tool air
