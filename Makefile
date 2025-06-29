templ:
	templ generate

tailwind-clean:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --clean

# Run tailwindcss to generate the styles.css bundle in watch mode.
tailwind-watch:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

dev:
	air \
	--build.cmd "tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --minify && go generate ./... && go build -o tmp/bin/main ./cmd/api/main.go" \
	--build.bin "tmp/bin/main" \
	--build.delay "100" \
	--build.exclude_dir "node_modules,assets,tmp" \
	--build.exclude_regex ".*_templ.go" \
	--build.include_ext "go,templ" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true \
	--proxy.enabled true \
	--proxy.proxy_port 3000 \
	--proxy.app_port 8080
