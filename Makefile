generate: generate/templ

generate/templ:
	@go tool templ generate

watch/build: generate
	@go build -o ./tmp/server ./cmd/server

watch:
	@go tool air
