
.PHONY: generate gen-back gen-front up build build-dev

generate: gen-back gen-front

gen-back:
	@WG_CONFIG_PATH=tmp/config/config.local.yaml GOWORK=off go generate ./...

gen-front:
	@docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-g typescript-fetch \
		-i /local/ent/openapi.json \
		-o /local/frontend/src/lib/api \
		--additional-properties=addResponseHeaders=true \
		--parameter-name-mappings x-page=page,x-items-per-page=itemsPerPage

build-dev:
	@docker build -t wg:dev . -f build/dev.Dockerfile

build:
	@docker build -t wg:latest . -f build/Dockerfile

up:
	@cd build && docker-compose up -d
