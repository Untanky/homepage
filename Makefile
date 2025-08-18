clean:
	rm -rf dist

tailwind-build:
	npx @tailwindcss/cli -i cmd/assets/index.css -o dist/assets/index.css

templ-generate:
	templ generate

go-build: clean tailwind-build templ-generate
	go build -o dist/server ./...

dev: go-build
	cd dist && ./server

