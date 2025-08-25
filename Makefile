clean:
	rm -rf dist

tailwind-build:
	npx @tailwindcss/cli -i cmd/assets/index.css -o dist/assets/index.css

templ-generate:
	templ generate

copy-migrations:
	cp -r db dist

go-build: clean tailwind-build templ-generate copy-migrations
	go build -o dist/server ./cmd

dev: go-build
	cd dist && ./server

