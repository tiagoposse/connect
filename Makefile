
.PHONY: generate gen-back gen-front up build build-dev

generate: gen-back gen-front backend-dev backend bin up

gen-back:
	@WG_CONFIG_PATH=tmp/config/config.local.yaml GOWORK=off go generate ./...

gen-front:
	@docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-g typescript-fetch \
		-i /local/ent/openapi.json \
		-o /local/frontend/src/lib/api \
		--additional-properties=addResponseHeaders=true,paramNaming=camelCase \
		--parameter-name-mappings x-page=page,x-items-per-page=itemsPerPage

backend-dev:
	@docker build -t wg:dev . -f build/dev.Dockerfile

backend:
	@docker build -t wg:latest . -f build/Dockerfile

bin:
	@CGO_ENABLED=0 go build -o bin/wg

up:
	@cd build && docker-compose up -d

seed:
	@cd build && ./init-admin.sh

cert:
	@openssl req -new -x509 -sha256 -key tmp/server.key -out tmp/server.crt -days 3650
