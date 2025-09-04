clean:
	rm -rf dist internal/assets/dist

assets-build:
	mkdir dist
	node esbuild.mjs

templ-generate:
	templ generate

copy-migrations:
	cp -r db dist

go-build: clean assets-build templ-generate copy-migrations
	go build -o dist/server ./cmd

dev: go-build
	cd dist && ./server

