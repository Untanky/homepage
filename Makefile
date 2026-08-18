generate: generate/templ

generate/templ:
	@go tool templ generate

build/web:
	@node ./esbuild.ts

watch/build: generate build/web
	@go build -o ./tmp/homepage ./cmd/homepage

watch:
	@go tool air

migrate/create:
	@read -p "Enter migration name: " MIGRATION_NAME; \
	go tool migrate create -seq -digits 6 -dir db/migrations -ext sql "$$MIGRATION_NAME"

migrate:
	@go run ./cmd/homepage migrations up
